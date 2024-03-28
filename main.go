package main

import (
	"fmt"
	"ticket-booking/helper"
)

type UserDetails struct {
	firstName     string
	lastName      string
	userEmail     string
	bookedTickets uint
}

func main() {

	x, y := swap("Hello", "World")
	fmt.Print(x, y)

	var firstName string
	var lastName string
	var userEmail string = "random@test.com"
	var bookedTickets uint
	var remainingTickets uint = 3

	bookingNames := make([]UserDetails, 0)

	titleName := "Ticket Booking Platform"

	for {
		firstNames := []string{}
		helper.GreetUsers(titleName)

		if remainingTickets > 0 {
			fmt.Printf("We have %v tickets remaining.\n", remainingTickets)

			// First Name
			fmt.Print("Enter your first name: ")
			fmt.Scan(&firstName)

			// Last Name
			fmt.Print("Enter your last name: ")
			fmt.Scan(&lastName)

			// Tickets
			fmt.Print("Enter number of tickets: ")
			fmt.Scan(&bookedTickets)

			fmt.Println()

			if helper.IsEntriesValid(firstName, lastName, userEmail, bookedTickets, remainingTickets) {

				var userData = UserDetails{
					firstName:     firstName,
					lastName:      lastName,
					userEmail:     userEmail,
					bookedTickets: bookedTickets,
				}

				fmt.Printf("%v booked %v Ticket(s).🎉\n", userData.firstName, bookedTickets)
				fmt.Printf("Confirmation will be send at %v\n", userEmail)
				remainingTickets -= bookedTickets

				bookingNames = append(bookingNames, userData)

				for _, booking := range bookingNames {
					firstNames = append(firstNames, booking.firstName)
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
