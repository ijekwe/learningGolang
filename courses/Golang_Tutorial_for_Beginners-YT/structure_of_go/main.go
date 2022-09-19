package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remaingingTickets int = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remaingingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available.\n", conferenceTickets, remaingingTickets)
	fmt.Println("Get your tickets here to attend!")

	var userName string
	var userTickets int
	// ask user for thier name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&userName)

	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
