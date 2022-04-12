package main

import (
	"fmt"
	"os"

	"github.com/yonikosiner/python-deez-nuts-with-sea-men/pkg/members"
	"github.com/yonikosiner/python-deez-nuts-with-sea-men/pkg/sponsor"
	"github.com/yonikosiner/python-deez-nuts-with-sea-men/pkg/user"
	"github.com/yonikosiner/python-deez-nuts-with-sea-men/utils"
)

func main() {
	var u *user.User
	var m *members.Members
	var s *sponsor.Sponser

	for {
		input := utils.AskForInput(`S to sign up
V to see all memmbers who are working as volunteers
E for members who are volunteering at the entrance gate
G for members who are volunteering at the gift shop
P for members who are helping with painting and dectoring
$ For members who have not yet paid there $75 feeop re
£ For members who subscription has expried and not yet paied for the year
N To become a new sponser, and also learn more about becoming a sponser
L to list all sponsers
C to exit`)

		//If we don't add the trim to remove the last charcetr we get the
		//charcter a new line feed (10 in ascii), and this does not let the
		//switch case work
		switch input[:len(input)-1] {
		case "S":
			user := u.NewUser()
			u.SaveUser(user)
		case "V":
			fmt.Println(m.GetVolunteeringMembers())
		case "E":
			fmt.Println(m.GetAreaOfVolunteer(user.EntranceGate))
		case "G":
			fmt.Println(m.GetAreaOfVolunteer(user.GiftShop))
		case "P":
			fmt.Println(m.GetAreaOfVolunteer(user.PaintingDecorating))
		case "$":
			fmt.Println(m.GetNonPayedMembers())
		case "£":
			fmt.Println(m.SubscriptionExpired())
		case "N":
			sponser := s.NewSponser()
			s.SaveSponsor(sponser)
		case "L":
			fmt.Println(s.GetSponers())
		case "C":
			os.Exit(0)
		}
	}
}
