package user

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type userImpl struct {
	db *sqlx.DB
}

func (service *userImpl) Create(data UserData) error {
	err := service.validateUser(data)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	err = service.addUser(data)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (service *userImpl) Mark(username string, mark Mark) error {
	err := service.validateMark(mark)
	if err != nil {
		return fmt.Errorf("failed to mark user: %w", err)
	}

	err = service.setMark(username, mark)
	if err != nil {
		return fmt.Errorf("failed to mark user: %w", err)
	}

	return nil
}

func (service *userImpl) GetInfo(username string) (UserData, error) {
	user, err := service.getUser(username)
	if err != nil {
		return UserData{}, fmt.Errorf("failed to get info: %w", err)
	}

	return user, nil
}

func (service *userImpl) GetInfos() ([]UserData, error) {
	users, err := service.getUsers()
	if err != nil {
		return []UserData{}, fmt.Errorf("failed to get infos: %w", err)
	}

	return users, nil
}

func (service *userImpl) validateUser(data UserData) error {
	return nil // TODO: Implement
}

func (service *userImpl) addUser(data UserData) error {
	return nil // TODO: Implement
}

func (service *userImpl) validateMark(mark Mark) error {
	return nil // TODO: Implement
}

func (service *userImpl) setMark(username string, mark Mark) error {
	return nil // TODO: Implement
}

func (service *userImpl) getUser(username string) (UserData, error) {
	return UserData{}, nil // TODO: Implement
}

func (service *userImpl) getUsers() ([]UserData, error) {
	return []UserData{}, nil // TODO: Implement
}
