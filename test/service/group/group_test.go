package group_test

import (
	"reflect"
	"testing"

	"github.com/acm-uiuc/core/model"
	"github.com/acm-uiuc/core/service/group"
)

func TestGetGroups(t *testing.T) {
	svc, err := group.New()
	if err != nil {
		t.Fatal(err)
	}

	groups, err := svc.GetGroups()
	if err != nil {
		t.Fatal(err)
	}

	expectedGroupTypeCount := 1
	if len(groups) != expectedGroupTypeCount {
		t.Fatalf("expected '%d', got '%d'", expectedGroupTypeCount, len(groups))
	}

	committees, ok := groups[model.GroupCommittees]
	if !ok {
		t.Fatalf("couldn't find group type: %s", model.GroupCommittees)
	}

	expectedCommitteeCount := 1
	if len(committees) != expectedCommitteeCount {
		t.Fatalf("expected '%d', got '%d'", expectedCommitteeCount, len(committees))
	}

	expected := model.Group{
		Name:        "Top4",
		Description: "The leadership team for ACM",
		Chairs:      "Arnav Sankaran",
		Members: []model.GroupMember{
			model.GroupMember{
				Role:        "Chair",
				Username:    "arnavs3",
				DisplayName: "Arnav Sankaran",
				Email:       "acm@illinois.edu",
			},
			model.GroupMember{
				Role:        "Vice Chair",
				Username:    "asdale2",
				DisplayName: "Asher Dale",
				Email:       "vice-chair@acm.illinois.edu",
			},
			model.GroupMember{
				Role:        "Treasurer",
				Username:    "mpj4",
				DisplayName: "Martin Juskelis",
				Email:       "treasurer@acm.illinois.edu",
			},
			model.GroupMember{
				Role:        "Secretary",
				Username:    "devyesh2",
				DisplayName: "Dev Satpathy",
				Email:       "secretary@acm.illinois.edu",
			},
		},
	}

	if !reflect.DeepEqual(expected, committees[0]) {
		t.Fatalf("expected '%+v', got '%+v'", expected, committees[0])
	}
}
