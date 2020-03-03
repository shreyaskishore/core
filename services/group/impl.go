package group

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

const (
	dataTTL = 15

	// TODO: Update this to use the real group store in the future
	dataURI = "https://gist.githubusercontent.com/ASankaran/a8f36ebb498a2098a9d49d5fbaf530cd/raw/932e382783b3bfe0fcc65937a7e2a35b1d6de128/groups.yaml"
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
	resp, err := http.Get(dataURI)
	if err != nil {
		return fmt.Errorf("failed to retrieve data: %w", err)
	}

	defer resp.Body.Close()

	rawData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read raw data: %w", err)
	}

	data := groupData {}
	err = yaml.Unmarshal(rawData, &data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal raw data: %w", err)
	}

	service.data = data
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