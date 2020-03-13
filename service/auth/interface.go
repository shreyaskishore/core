package auth

import (
	"github.com/acm-uiuc/core/model"
)

type AuthService interface {
	GetOAuthRedirect(provider string) (string, error)
	Authorize(provider string, code string) (*model.Token, error)
	Verify(token string) (bool, string, error)
}

func New() (AuthService, error) {
	return nil, nil
}
