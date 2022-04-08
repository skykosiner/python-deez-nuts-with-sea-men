package user

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/yonikosiner/python-deez-nuts-with-sea-men/utils"
)

func areaToVolunteer(volunteer bool) AreasToVolunteer {
    if volunteer {
    }

    return None
}


func (u *User) NewUser() bool {
    firstName := utils.AskForInput("Enter your first name? ")
    lastName := utils.AskForInput("Enter your last name? ")

    paying := utils.StringToBoolean(utils.AskForInput("Would you like to pay now? (yes/no) "))
    volunteer := utils.StringToBoolean(utils.AskForInput("Would you like to volunteer? (yes/no) "))

    fullname := &FullName{firstName, lastName}
    user := &User{*fullname, utils.GetCurrentDate(), paying, volunteer, areaToVolunteer(volunteer)}

    fmt.Println(user)

    user.saveUser(*user)
    fmt.Print(user.getAllUsers())

    return true
}

func (u *User) saveUser(user User) error {
    b, err := json.Marshal(user)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(string(b))

    return os.WriteFile("./users.json", []byte(b), 0600)
}

func (u *User) getAllUsers() []User {
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

    var data []User

    for _, value := range text {
        fmt.Println(value)
        err := json.Unmarshal([]byte(value), &data)

        if err != nil {
            log.Fatal(err)
        }
    }

    return data
}

func (u *User) checkUserDup(user User) bool {
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

    return true
}
