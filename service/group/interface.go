package group

import (
	"fmt"

	"github.com/acm-uiuc/core/model"
)

type GroupService interface {
	GetGroups() (map[string][]model.Group, error)
	VerifyMembership(username string, groupType string, groupName string) (bool, error)
}

func New() (GroupService, error) {
	service := &groupImpl{
		lastUpdated: 0,
	}

	err := service.refreshData()
	if err != nil {
		return nil, fmt.Errorf("failed to create group service: %w", err)
	}

	return service, nil
}
