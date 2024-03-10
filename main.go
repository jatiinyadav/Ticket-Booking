package main

import (
	"fmt"
	"strings"
)

func main() {
	var firstName string
	var lastName string
	var fullName string
	var userEmail string = "random@test.com"
	var bookedTickets uint
	var remainingTickets uint = 3

	bookingNames := []string{}

	titleName := "Ticket Booking Platform"

	for {
		firstNames := []string{}
		greetUsers(titleName)

		if remainingTickets > 0 {
			fmt.Printf("We have %v tickets remaining.\n", remainingTickets)

			// First Name
			fmt.Print("Enter your first name: ")
			fmt.Scan(&firstName)

			// Last Name
			fmt.Print("Enter your last name: ")
			fmt.Scan(&lastName)

			fullName = firstName + " " + lastName
			bookingNames = append(bookingNames, fullName)

			// Tickets
			fmt.Print("Enter number of tickets: ")
			fmt.Scan(&bookedTickets)

			fmt.Println()

			if isEntriesValid(firstName, lastName, userEmail, bookedTickets, remainingTickets) {
				fmt.Printf("%v booked %v Ticket(s).🎉\n", fullName, bookedTickets)
				fmt.Printf("Confirmation will be send at %v\n", userEmail)
				remainingTickets -= bookedTickets

				for _, booking := range bookingNames {
					firstNames = append(firstNames, strings.Fields(booking)[0])
				}

				fmt.Printf("All Users: %v\n", firstNames)

			} else {
				fmt.Print("Data is invalid, try again.")
				// fmt.Printf("You cannot book %v tickets, we have %v left\n", bookedTickets, remainingTickets)
				continue
			}

		} else {
			fmt.Printf("Sorry, show is full😔.\n")
			return
		}
	}

}

func greetUsers(titleName string) {
	fmt.Println()
	fmt.Printf("Welcome to %v ❤️ \n", titleName)
}

func isEntriesValid(firstName string, lastName string, userEmail string, bookedTickets uint, remainingTickets uint) bool {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTicket := bookedTickets > 0 && bookedTickets <= remainingTickets

	return isValidName && isValidEmail && isValidTicket
}
