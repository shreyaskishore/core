package user

import (
	"github.com/jmoiron/sqlx"
)

type userImpl struct {
	db *sqlx.DB
}

func (service *userImpl) Create(data UserData) error {
	return nil
}

func (service *userImpl) Mark(username string, mark string) error {
	return nil
}

func (service *userImpl) GetInfo(username string) (UserData, error) {
	return UserData{}, nil
}

func (service *userImpl) GetInfos() ([]UserData, error) {
	return []UserData{}, nil
}
