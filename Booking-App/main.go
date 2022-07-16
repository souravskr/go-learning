package main

import "fmt"

func main() {
	const totalTickets = 50
	availableTickets := 50

	conferenceName := "Go Conference"

	// Welcome Message
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v, and %v tickets are available \n", totalTickets, availableTickets)
	fmt.Println("Get your tickets here to attend")

	// Booking Information
	var participantsNames []string
	// Take user input
	var userName string
	var email string
	var userTickets int
	wannaBook := "yes"

	for {
		fmt.Println("Do you want to book tickets? (yes/no)")
		fmt.Print("--> ")
		fmt.Scan(&wannaBook)
		if wannaBook == "no" {
			break
		} else if availableTickets <= 0 {
			fmt.Println("Tickets are booked out!")
			break
		}

		fmt.Println("ENTER your full name: ")
		fmt.Print("--> ")
		fmt.Scan(&userName)

		fmt.Println("ENTER your email: ")
		fmt.Print("--> ")
		fmt.Scan(&email)

		fmt.Println("ENTER number of tickets: ")
		fmt.Print("--> ")
		fmt.Scan(&userTickets)
		if userTickets > 25 {
			fmt.Println("You can't book more than 25 tickets")
			break
		}

		// Confirmation of booking tickets
		fmt.Printf("Thank you %v for booking %v tickets. You you will receive a confirmation email shortly \n", userName, userTickets)

		participantsNames = append(participantsNames, userName)
		availableTickets = availableTickets - userTickets
	}
	fmt.Printf("User list: %v and remaining tickets are %v\n", participantsNames, availableTickets)

}
