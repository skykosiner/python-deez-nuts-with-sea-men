package members

import (
	"github.com/yonikosiner/python-deez-nuts-with-sea-men/pkg/user"
	"github.com/yonikosiner/python-deez-nuts-with-sea-men/utils"
)

type Members struct {
    Users []*user.User
}

func (m *Members) getMembers() []user.User {
    var u *user.User
    users := u.GetAllUsers()
    return users
}

func (m *Members) getNonPayedMembers() []string {
    members := m.getMembers()
    var notPaying []string

    for _, value := range members {
        if !value.Paying {
            notPaying = append(notPaying, utils.GetFullName(value))
        }
    }

    return notPaying
}

func (m *Members) getVolunteeringMembers() []string {
    members := m.getMembers()
    var volunteers []string

    for _, value := range members {
        if value.Volunteer {
            volunteers = append(volunteers, utils.GetFullName(value))
        }
    }

    return volunteers
}

// This works with any of the three areas, you just pass in
// areasToVolunteer.painting_decorating or any other area one can volunteer
func (m *Members) getAreaOfVolunteer(area user.AreasToVolunteer) []string {
    members := m.getMembers()
    var volunteers []string

    for _, value := range members {
        if value.AreaToVolunteer == area {
            volunteers = append(volunteers, utils.GetFullName(value))
        }
    }

    return volunteers
}
