package sponsor

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/yonikosiner/python-deez-nuts-with-sea-men/pkg/user"
	"github.com/yonikosiner/python-deez-nuts-with-sea-men/utils"
)

func (s *Sponser) NewSponser() Sponser {
	fmt.Println("Have your name and a message of your choise be put on a brass plaque, and also support as with a $200 fee <3")

	firstName := utils.AskForInput("Enter your first name? ")
	lastName := utils.AskForInput("Enter your last name? ")
	fullname := user.FullName{firstName, lastName}

	var message string

	message = utils.AskForInput("What would you like your plank to say?")

	if len(message) > 64 {
		newMessage := utils.StringToBoolean(utils.AskForInput("Your message is to long, would you like to enter a new message? (yes/no) "))
		if newMessage {
			message = utils.AskForInput("Entere in your new message ")
		}
	}

	sponser := s.checkSponsorDup(Sponser{fullname, message})

	return sponser
}

func (s *Sponser) checkSponsorDup(sponser Sponser) Sponser {
	sponsers := s.getAllSponsors()

	for _, value := range sponsers {
		if sponser.Name == value.Name {
			log.Fatal("You're already sponsoring")
			contiuneSponser := utils.StringToBoolean(utils.AskForInput("Continue as someone else? (yes/no) "))

			if contiuneSponser {
				return s.NewSponser()
			}
		}
	}
	return sponser
}

func (s *Sponser) SaveSponsor(sponser Sponser) {
	b, err := json.Marshal(sponser)

	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile("./sponsers.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)

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

func (s *Sponser) GetSponers() []string {
	sponsers := s.getAllSponsors()

	var sponsersReturn []string

	for _, value := range sponsers {
		name := fmt.Sprintf("Name: %s %s\n", value.Name.FirstName, value.Name.LastName)
		message := fmt.Sprintf("Message: %s", value.Message)

		stringReturn := name + message
		sponsersReturn = append(sponsersReturn, stringReturn)
	}

	return sponsersReturn
}

func (s *Sponser) getAllSponsors() []Sponser {
	b, err := os.Open("./sponsers.json")

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

	var sponsers []Sponser

	for _, value := range text {
		var data Sponser
		err := json.Unmarshal([]byte(value), &data)

		if err != nil {
			log.Fatal(err)
		}

		sponsers = append(sponsers, data)
	}

	return sponsers
}
