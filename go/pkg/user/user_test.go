package user

import (
	"testing"

	"github.com/yonikosiner/python-deez-nuts-with-sea-men/utils"
	// "github.com/yonikosiner/python-deez-nuts-with-sea-men/utils"
)

func makeDumbyUser() User {
	firstName := "Yoni"
	lastName := "Kosiner"

	paying := true
	volunteer := true

	fullname := &FullName{firstName, lastName}

	return User{*fullname, utils.GetCurrentDate(), paying, volunteer, EntranceGate, SetSubcription(volunteer)}
}

func TestAreaToVolunteerTrue(t *testing.T) {
	results := getAreaToVolunteer(true)

	if results != EntranceGate && results != GiftShop && results != PaintingDecorating {
		t.Errorf("Fool of a took, you're not one of the three results, got: %s", string(rune(results)))
	}
}

func TestAreaToVolunteerFalse(t *testing.T) {
	results := getAreaToVolunteer(false)

	if results != None {
		t.Errorf("Expected None, got %s", string(rune(results)))
	}
}

/* func TestGetAllUsers(t *testing.T) {
	var u *User

	results := u.GetAllUsers()
} */

func TestCheckUserDupFalse(t *testing.T) {
	var u *User

	user := makeDumbyUser()

	results := u.checkUserDup(user)

	if results != user {
		t.Errorf("The user is not a duplicate (even though it should be), got the object %s", string(results.Name.FirstName))
	}
}
