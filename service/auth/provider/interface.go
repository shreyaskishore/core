package provider

import (
	"fmt"
)

type OAuthProvider interface {
	GetOAuthRedirect() (string, error)
	GetOAuthToken(code string) (string, error)
	GetVerifiedEmail(token string) (string, error)
}

func GetProvider(provider string) (OAuthProvider, error) {
	return nil, fmt.Errorf("not implemented")
}
