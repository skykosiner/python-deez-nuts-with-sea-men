package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func AskForInput(message string) string {
	var stringToReturn string

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("%s \n", message)
		text, _ := reader.ReadString('\n')

		valid := validateInput(text)

		stringToReturn = text

		if valid {
			break
		}
	}
	return stringToReturn
}

func validateInput(input string) bool {
	// Remove the weird 10 at the end of the text object smh
	input = input[:len(input)-1]

	// Make sure that there is no weird ascii charters
	for i := 0; i < len(input); i++ {
		byteValue := input[i]
		intValue := int(byteValue)

		if intValue < 32 || intValue > 127 {
			return false
		}
	}

	return true
}

func GetCurrentDate() string {
	return time.Now().Format("02-01-06")
}

func StringToBoolean(value string) bool {
	// Remove ascii 10 (new line feed) at the end of the line so that the case
	// statement works
	value = value[:len(value)-1]

	switch value {
	case "yes", "y":
		return true
	case "no", "n":
		return false
	default:
		log.Fatal("Sorry there was an error, please try again: StringToBoolean")
		return false
	}
}
