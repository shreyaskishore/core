package group

import (
	"fmt"
	"time"
)

const (
	dataTTL = 15
)

type groupImpl struct {
	data groupData
	lastUpdated int64
}

type groupData struct {
	groups map[string][]Group
}

func (service *groupImpl) GetMemberships(username string) ([]string, error) {
	if service.isDataStale() {
		err := service.refreshData();
		if err != nil {
			return []string{}, fmt.Errorf("failed to get memberships: %w", err)
		}
	}

	memberships := []string{}
	for _, groups := range service.data.groups {
		for _, group := range groups {
			for _, member := range group.Members {
				if member.Username == username {
					memberships = append(memberships, group.Name)
				}
			}
		}
	}

	return memberships, nil
}

func (service *groupImpl) GetGroups(groupType GroupType) ([]Group, error) {
	if service.isDataStale() {
		err := service.refreshData();
		if err != nil {
			return []Group{}, fmt.Errorf("failed to get groups: %w", err)
		}
	}

	err := service.validateGroupType(groupType)
	if err != nil {
		return []Group{}, fmt.Errorf("failed to get groups: %w", err)
	}

	groups, ok := service.data.groups[string(groupType)]
	if !ok {
		return []Group{}, fmt.Errorf("failed to find data for group type: %w", err)
	}

	return groups, nil
}

func (service *groupImpl) Verify(username string, groupType GroupType, groupName string) (bool, error) {
	if service.isDataStale() {
		err := service.refreshData();
		if err != nil {
			return false, fmt.Errorf("failed to verify membership: %w", err)
		}
	}

	groups, ok := service.data.groups[string(groupType)]
	if !ok {
		return false, fmt.Errorf("failed to find data for group type: %s", string(groupType))
	}

	for _, group := range groups {
		if group.Name == groupName {
			for _, member := range group.Members {
				if member.Username == username {
					return true, nil
				}
			}

			return false, nil
		}
	}

	return false, fmt.Errorf("failed to find data for group name: %s", groupName)
}

func (service *groupImpl) refreshData() error {
	// TODO: Pull the group data from github and parse into a groupData
	service.lastUpdated = time.Now().Unix()

	return nil
}

func (service *groupImpl) isDataStale() bool {
	return service.lastUpdated < time.Now().Add(-1 * dataTTL * time.Minute).Unix()
}

func (service *groupImpl) validateGroupType(groupType GroupType) error {
	for _, validGroupType := range validGroupTypes {
		if groupType == validGroupType {
			return nil
		}
	}

	return fmt.Errorf("invalid group type: %s", groupType)
}