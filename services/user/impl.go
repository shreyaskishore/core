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
	// TODO: Implement user data validate
	// For now everything is treated as valid
	return nil
}

func (service *userImpl) addUser(data UserData) error {
	_, err := service.db.NamedExec("INSERT INTO users (username, first_name, last_name, graduation_year, major, mark) VALUES (:username, :first_name, :last_name, :graduation_year, :major, :mark)", data)
	if err != nil {
		fmt.Errorf("failed to add user to database: %w", err)
	}

	return nil
}

func (service *userImpl) validateMark(mark Mark) error {
	for _, validMark := range validMarks {
		if mark == validMark {
			return nil
		}
	}

	return fmt.Errorf("invalid mark: %s", mark)
}

func (service *userImpl) setMark(username string, mark Mark) error {
	params := UserData {
		Username: username,
		Mark: string(mark),
	}

	_, err := service.db.NamedExec("UPDATE users SET mark=:mark WHERE username=:username", params)
	if err != nil {
		return fmt.Errorf("failed to update mark: %w", err)
	}

	return nil
}

func (service *userImpl) getUser(username string) (UserData, error) {
	params := UserData {
		Username: username,
	}

	rows, err := service.db.NamedQuery("SELECT * FROM users WHERE username=:username", params)
	if err != nil {
		return UserData{}, fmt.Errorf("failed to query database for user: %w", err)
	}

	result := UserData {}
	for rows.Next() {
		err := rows.StructScan(&result)
		if err != nil {
			return UserData{}, fmt.Errorf("failed to decode row from database: %w", err)
		}
	}

	err = rows.Err()
	if err != nil {
		return UserData{}, fmt.Errorf("failed reading rows from database: %w", err)
	}

	return result, nil
}

func (service *userImpl) getUsers() ([]UserData, error) {
	rows, err := service.db.NamedQuery("SELECT * FROM users", struct{}{})
	if err != nil {
		return []UserData{}, fmt.Errorf("failed to query database for users: %w", err)
	}

	results := []UserData {}
	for rows.Next() {
		result := UserData {}
		err := rows.StructScan(&result)
		if err != nil {
			return []UserData{}, fmt.Errorf("failed to decode row from database: %w", err)
		}
		results = append(results, result)
	}

	err = rows.Err()
	if err != nil {
		return []UserData{}, fmt.Errorf("failed reading rows from database: %w", err)
	}

	return results, nil
}
