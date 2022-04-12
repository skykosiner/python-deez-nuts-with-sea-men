package user

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/yonikosiner/python-deez-nuts-with-sea-men/utils"
)

func getCurrentDateObject() Date {
	currentTime := time.Now()
	month := int(currentTime.Month())
	day := currentTime.Day()
	year := currentTime.Year()

	return Date{month, day, year}
}

func GetFullName(user User) string {
	return fmt.Sprintf("%s %s", user.Name.FirstName, user.Name.LastName)
}

func SetSubcription(volunteer bool) Subscription {
	if volunteer {
		return Subscription{getCurrentDateObject()}
	}

	return Subscription{Date{0, 0, 0}}
}

func areaToVolunteer(volunteer bool) AreasToVolunteer {
	if volunteer {
		areas := [3]string{"EntranceGate", "GiftShop", "PaintingDecorating"}

		rand.Seed(time.Now().UnixNano())

		chosen := areas[rand.Intn(len(areas))]

		switch chosen {
		case "EntranceGate":
			return EntranceGate
		case "GiftShop":
			return GiftShop
		case "PaintingDecorating":
			return PaintingDecorating
		default:
			log.Fatal("Sorry there was an error, please try again: areaToVolunteer")
		}
	}

	return None
}

func (u *User) NewUser() User {
	firstName := utils.AskForInput("Enter your first name? ")
	lastName := utils.AskForInput("Enter your last name? ")

	paying := utils.AskForInput("Would you like to pay now? (yes/no) ")
	payingBool := utils.StringToBoolean(paying)

	volunteer := utils.AskForInput("Would you like to volunteer? (yes/no) ")
	volunteerBool := utils.StringToBoolean(volunteer)

	fullname := &FullName{firstName, lastName}

	user := &User{*fullname, utils.GetCurrentDate(), payingBool, volunteerBool, areaToVolunteer(volunteerBool), SetSubcription(volunteerBool)}

	return u.checkUserDup(*user)
}

func (u *User) SaveUser(user User) {
	b, err := json.Marshal(user)

	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile("./users.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		log.Fatal(err)
	}

	if _, err := f.Write([]byte(fmt.Sprintf("%s\n", b))); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func (u *User) GetAllUsers() []User {
	b, err := os.Open("./users.json")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(b)
	scanner.Split(bufio.ScanLines)

	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	b.Close()

	var users []User

	for _, value := range text {
		var data User
		err := json.Unmarshal([]byte(value), &data)

		if err != nil {
			log.Fatal(err)
		}

		users = append(users, data)
	}

	return users
}

func (u *User) checkUserDup(user User) User {
	users := u.GetAllUsers()

	for _, value := range users {
		if user.Name == value.Name {
			log.Fatal("Sorry that user is already in our database")
			return u.NewUser()
		}
	}

	return user
}

func (u *User) SubscriptionExpired() bool {
	// Make sure that the pay is not set to 0
	if u.Subscription.LastPaid.Day != 0 {
		subscirption := u.Subscription
		// Weird that the date is formated in 2006 not 1970 (the start of time, well at least it is to computers)
		today := time.Now()
		if subscirption.LastPaid.Month < int(today.Month()) && subscirption.LastPaid.Day < today.Day() && subscirption.LastPaid.Year < today.Year() {
			return true
		} else {
			return false
		}
	}

	return false
}
