package group

import (
	"github.com/acm-uiuc/core/model"
)

type GroupService interface {
	GetGroups() ([]model.Group, error)
	VerifyMembership(username string, group string) (bool, error)
}

func New() (GroupService, error) {
	return nil, nil
}
