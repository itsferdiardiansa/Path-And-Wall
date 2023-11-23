package pointer

import "fmt"

func RunPointer() {
	var firstName string
	var lastName string
	var email string
	var userTickets string

	fmt.Println("Enter your firstname: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your lastname: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter your ticket amount: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v ticket(s). The confirmation order will be sent to your email %v.\n", firstName, lastName, userTickets, email)
}
