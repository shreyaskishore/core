package group

type groupImpl struct {

}

func (service *groupImpl) GetMemberships(username string) ([]string, error) {
	return []string{}, nil
}

func (service *groupImpl) GetGroups(groupType string) ([]Group, error) {
	return []Group{}, nil
}

func (service *groupImpl)  Verify(username string, groupName string) (bool, error) {
	return false, nil
}
