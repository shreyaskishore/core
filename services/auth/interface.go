package auth

import (
	"fmt"

	"github.com/acm-uiuc/core/database"
)

type AuthService interface {
	Login(username string, password string) (string, error)
	Logout(token string) error
	Verify(token string) (string, error)
	CreateLocalAccount(username string, password string) error
}

func New() (AuthService, error) {
	db, err := database.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create auth service: %w", err)
	}

	return &authImpl{
		db: db,
	}, nil
}
