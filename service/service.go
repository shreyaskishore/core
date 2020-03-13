package service

import (
	"fmt"

	"github.com/acm-uiuc/core/service/auth"
	"github.com/acm-uiuc/core/service/group"
	"github.com/acm-uiuc/core/service/user"
)

type Service struct {
	Auth  auth.AuthService
	User  user.UserService
	Group group.GroupService
}

func New() (*Service, error) {
	authService, err := auth.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create auth service: %w", err)
	}

	userService, err := user.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create user service: %w", err)
	}

	groupService, err := group.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create group service: %w", err)
	}

	return &Service{
		Auth:  authService,
		User:  userService,
		Group: groupService,
	}, nil
}
