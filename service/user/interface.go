package user

import (
	"fmt"

	"github.com/acm-uiuc/core/database"
	"github.com/acm-uiuc/core/model"
)

type UserService interface {
	GetUser(username string) (*model.User, error)
	CreateUser(user model.User) error
	GetUsers() ([]model.User, error)
	MarkUser(username string, mark string) error
}

func New() (UserService, error) {
	db, err := database.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create user service: %w", err)
	}

	return &userImpl{
		db: db,
	}, nil
}
