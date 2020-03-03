package services

import (
	"fmt"

	"github.com/acm-uiuc/core/services/auth"
	"github.com/acm-uiuc/core/services/user"
	"github.com/acm-uiuc/core/services/group"
)

type Services struct {
	Auth auth.AuthService
	User user.UserService
	Group group.GroupService
}

func New() (Services, error) {
	auth, err := auth.New()
	if err != nil {
		return Services{}, fmt.Errorf("failed to create auth service: %w", err)
	}

	user, err := user.New()
	if err != nil {
		return Services{}, fmt.Errorf("failed to create user service: %w", err)
	}

	group, err := group.New()
	if err != nil {
		return Services{}, fmt.Errorf("failed to create group service: %w", err)
	}

	return Services{
		Auth: auth,
		User: user,
		Group: group,
	}, nil
}
