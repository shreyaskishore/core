package group

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/acm-uiuc/core/config"
	"github.com/acm-uiuc/core/model"
)

const (
	dataTTL = 15
)

type groupImpl struct {
	data        map[string][]model.Group
	lastUpdated int64
}

func (service *groupImpl) GetGroups() (map[string][]model.Group, error) {
	err := service.handleStaleData()
	if err != nil {
		return nil, fmt.Errorf("failed to update stale data: %w", err)
	}

	return service.data, nil
}

func (service *groupImpl) VerifyMembership(username string, groupType string, groupName string) (bool, error) {
	err := service.handleStaleData()
	if err != nil {
		return false, fmt.Errorf("failed to update stale data: %w", err)
	}

	groups, ok := service.data[groupType]
	if !ok {
		return false, fmt.Errorf("failed to find data for group type: %s", groupType)
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
	uri, err := config.GetConfigValue("GROUP_URI")
	if err != nil {
		return fmt.Errorf("failed to get group data uri: %w", err)
	}

	resp, err := http.Get(uri)
	if err != nil {
		return fmt.Errorf("failed to retrieve data: %w", err)
	}

	defer resp.Body.Close()

	rawData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read raw data: %w", err)
	}

	err = yaml.Unmarshal(rawData, &service.data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal raw data: %w", err)
	}

	service.lastUpdated = time.Now().Unix()

	return nil
}

func (service *groupImpl) handleStaleData() error {
	if service.lastUpdated < time.Now().Add(-1*dataTTL*time.Minute).Unix() {
		return service.refreshData()
	}

	return nil
}
