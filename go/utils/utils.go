package utils

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func AskForInput(message string) string {
    stringToReturn := ""
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
    // wait you're not suposed to indent this? This just feels wrong, but hey
    // if the go people say so I guess
    switch value {
    case "yes", "y":
        return true
    case "no", "n":
        return false
    }

    return false
}
