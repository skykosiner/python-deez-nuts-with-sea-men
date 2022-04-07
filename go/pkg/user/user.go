package user

import (
	"github.com/yonikosiner/python-deez-nuts-with-sea-men/utils"
    "math/rand"
    "time"
)

type AreasToVolunteer int

const (
    None AreasToVolunteer = iota
    EntranceGate = 1
    GiftShop
    PaintingDecorating
)

type FullName struct {
    FirstName string
    LastName string
}

type User struct {
    Name FullName
    SignUpDate string
    Paying bool
    Volunteer bool
    AreaToVolunteer AreasToVolunteer
}


func (u *User) NewUser() bool {
    firstName := utils.AskForInput("Enter your first name? ")
    lastName := utils.AskForInput("Enter your last name? ")

    paying := utils.StringToBoolean(utils.AskForInput("Would you like to pay now? (yes/no) "))
    volunteer := utils.StringToBoolean(utils.AskForInput("Would you like to volunteer? (yes/no) "))

    // If the user choses to volunteer pick a random area for them to work
    if volunteer {
    }

    fullname := &FullName{firstName, lastName}
    user := &User{*fullname, utils.GetCurrentDate(), paying, volunteer}

    return true
}
