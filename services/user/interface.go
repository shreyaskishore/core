package user

import (
	"fmt"

	"github.com/acm-uiuc/core/database"
)

const (
	MarkBasic string = "BASIC"
	MarkPaid string = "PAID"
	MarkRecruiter string = "RECRUITER"
)

type UserData struct {
	Username string `db:"username"`
	FirstName string `db:"first_name"`
	LastName string `db:"last_name"`
	GraduationYear int32 `db:"graduation_year"`
	Major string `db:"major"`
	Mark string `db:"mark"`
}

type UserService interface {
	Create(data UserData) error
	Mark(username string, mark string) error
	GetInfo(username string) (UserData, error)
	GetInfos() ([]UserData, error)
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
