package group

import (
	"fmt"
)

type Group struct {
	Name string
	Description string
	Chairs string
	Members []GroupMember
	MeetingTime string
	MeetingLocation string
	Website string
	Email string
}

type GroupMember struct {
	Role string
	Username string
	DisplayName string
	Email string
}

type GroupType string
const (
	GroupCommittee GroupType = "committee"
	GroupSIG GroupType = "sig"
)
var validGroupTypes = []GroupType{GroupCommittee, GroupSIG}

type GroupService interface {
	GetMemberships(username string) ([]string, error)
	GetGroups(groupType GroupType) ([]Group, error)
	Verify(username string, groupType GroupType, groupName string) (bool, error)
}

func New() (GroupService, error) {
	service := &groupImpl {
		lastUpdated: 0,
	}
	
	err := service.refreshData()
	if err != nil {
		return nil, fmt.Errorf("failed to create group service: %w", err)
	}

	return service, nil
}
