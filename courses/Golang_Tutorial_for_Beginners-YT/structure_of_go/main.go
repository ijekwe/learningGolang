package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remaingingTickets = 50

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available.\n", conferenceTickets, remaingingTickets)
	fmt.Println("Get your tickets here to attend!")

}
