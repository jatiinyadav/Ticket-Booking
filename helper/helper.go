package helper

import (
	"fmt"
	"strings"
)

func GreetUsers(titleName string) {
	fmt.Println()
	fmt.Printf("Welcome to %v ❤️ \n", titleName)
}

func IsEntriesValid(firstName string, lastName string, userEmail string, bookedTickets uint, remainingTickets uint) bool {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTicket := bookedTickets > 0 && bookedTickets <= remainingTickets

	return isValidName && isValidEmail && isValidTicket
}
