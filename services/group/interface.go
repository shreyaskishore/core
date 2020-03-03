package group

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

type GroupService interface {
	GetMemberships(username string) ([]string, error)
	GetGroups(groupType string) ([]Group, error)
	Verify(username string, groupName string) (bool, error)
}

func New() (GroupService, error) {
	return &groupImpl{}, nil
}
