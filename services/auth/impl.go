package auth

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"time"

	ldap "github.com/jtblin/go-ldap-client"

	"github.com/jmoiron/sqlx"
)

const (
	ldapBase = "ou=People,dc=ad,dc=uillinois,dc=edu"
	ldapHost = "ldap.ad.uillinois.edu"
	ldapPort = 389

	tokenLength = 64
	tokenLifetime = 24
)

type authImpl struct {
	db *sqlx.DB
}

type tokenRow struct {
	Username string `db:"username"`
	Token string `db:"token"`
	Expiration int64 `db:"expiration"`
}

// TODO: Persist tokens in a sql table
// Stores token -> username
var tokens = map[string]string{}

func (service *authImpl) Login(username string, password string) (string, error) {
	// Start with ldap login and fall back to local login if needed
	loginStrategies := [](func(string, string) (bool, error)){service.ldapLogin, service.localLogin}

	var success bool
	var errs []error
	for _, loginStrategy := range loginStrategies {
		var err error
		success, err = loginStrategy(username, password)
		if err != nil {
			errs = append(errs, err)
		}
		if success {
			break
		}
	}

	if !success {
		collectedErrors := fmt.Errorf("")
		for _, err := range errs {
			collectedErrors = fmt.Errorf("%s | %w", collectedErrors.Error(), err)
		}
		return "", fmt.Errorf("login failed: %w", collectedErrors)
	}

	token := service.generateTokenString()

	err := service.addToken(token, username)
	if err != nil {
		return "", fmt.Errorf("failed to add token: %w", err)
	}

	return token, nil
}

func (service *authImpl) Logout(token string) error {
	err := service.removeToken(token)
	if err != nil {
		return fmt.Errorf("failed to remove token: %w", err)
	}

	return nil
}

func (service *authImpl) Verify(token string) (string, error) {
	username, err := service.tokenToUsername(token)
	if err != nil {
		return "", fmt.Errorf("verification failed: %w", err)
	}

	return username, nil
}

func (service *authImpl) CreateLocalAccount(username string, password string) error {
	return errors.New("not implemented")
}

func (service *authImpl) ldapLogin(username string, password string) (bool, error) {
	// We don't have a readonly service user, so we'll just use the account
	// we're authenticating. If it fails to bind, then we know the credentials
	// are invalid.
	bindDN := fmt.Sprintf("cn=%s,%s", username, ldapBase)

	client := &ldap.LDAPClient{
		Base: ldapBase,
		Host: ldapHost,
		Port: ldapPort,
		ServerName: ldapHost,
		UseSSL: false, // We're using TLS
		BindDN: bindDN,
		BindPassword: password,
		UserFilter: "(cn=%s)",
		Attributes: []string{},
	}

	defer client.Close()

	success, _, err := client.Authenticate(username, password)
	if err != nil {
		return false, fmt.Errorf("error in ldap authentication: %w", err)
	}
	if !success {
		return false, fmt.Errorf("invalid ldap credentials for %s", username)
	}

	return true, nil
}

func (service *authImpl) localLogin(username string, password string) (bool, error) {
	return false, errors.New("not implemented")
}

func (service *authImpl) generateTokenString() string {
	token := make([]byte, tokenLength)
	rand.Read(token)
	return base64.StdEncoding.EncodeToString(token)
}

func (service *authImpl) addToken(token string, username string) error {
	params := tokenRow {
		Username: username,
		Token: token,
		Expiration: time.Now().Add(tokenLifetime * time.Hour).Unix(),
	}
	
	_, err := service.db.NamedExec("INSERT INTO tokens (username, token, expiration) VALUES (:username, :token, :expiration)", params)	
	if err != nil {
		return fmt.Errorf("failed to add token to database: %w", err)
	}

	return nil
}

func (service *authImpl) removeToken(token string) error {
	params := tokenRow {
		Token: token,
	}

	_, err := service.db.NamedExec("DELETE FROM tokens WHERE token=:token", params)
	if err != nil {
		return fmt.Errorf("failed to remove token from database: %w", err)
	}
	
	return nil
}

func (service *authImpl) tokenToUsername(token string) (string, error) {
	params := tokenRow {
		Token: token,
	}
	
	rows, err := service.db.NamedQuery("SELECT username, token, expiration FROM tokens WHERE token=:token", params)
	if err != nil {
		return "", fmt.Errorf("failed to query database from token: %w", err)
	}

	result := tokenRow {}
	for rows.Next() {
		err := rows.StructScan(&result)
		if err != nil {
			return "", fmt.Errorf("failed to decode row from database: %w", err)
		}
	}

	err = rows.Err()
	if err != nil {
		return "", fmt.Errorf("failed reading rows from database: %w", err)
	}

	return result.Username, nil
}
