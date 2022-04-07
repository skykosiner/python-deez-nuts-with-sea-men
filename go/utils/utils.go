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
    return time.Now().Format("DD-MM-YY")
}

func StringToBoolean(value string) bool {
    if value == "yes" || value == "y" {
        return true
    } else if value == "no" || value == "n" {
        return false
    }
}
