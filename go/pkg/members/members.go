package members

import (
	"github.com/yonikosiner/python-deez-nuts-with-sea-men/pkg/user"
)

type Members struct {
	Users []*user.User
}

func (m *Members) GetMembers() []user.User {
	var u *user.User
	users := u.GetAllUsers()
	return users
}

func (m *Members) GetNonPayedMembers() []string {
	members := m.GetMembers()
	var notPaying []string

	for _, value := range members {
		if !value.Paying {
			notPaying = append(notPaying, user.GetFullName(value))
		}
	}

	return notPaying
}

func (m *Members) GetVolunteeringMembers() []string {
	members := m.GetMembers()
	var volunteers []string

	for _, value := range members {
		if value.Volunteer {
			volunteers = append(volunteers, user.GetFullName(value))
		}
	}

	return volunteers
}

// This works with any of the three areas, you just pass in
// areasToVolunteer.painting_decorating or any other area one can volunteer
func (m *Members) GetAreaOfVolunteer(area user.AreasToVolunteer) []string {
	members := m.GetMembers()
	var volunteers []string

	for _, value := range members {
		if value.AreaToVolunteer == area {
			volunteers = append(volunteers, user.GetFullName(value))
		}
	}

	return volunteers
}

func (m *Members) SubscriptionExpired() []string {
	members := m.GetMembers()
	var expried []string

	for _, value := range members {
		if value.SubscriptionExpired() {
			expried = append(expried, user.GetFullName(value))
		}
	}

    return expried
}
