package user

import (
	"github.com/acm-uiuc/core/model"
)

type UserService interface {
	GetUser(username string) (*model.User, error)
	CreateUser(user model.User) error
	GetUsers() ([]model.User, error)
	MarkUser(username string, mark string) error
}

func New() (UserService, error) {
	return nil, nil
}
