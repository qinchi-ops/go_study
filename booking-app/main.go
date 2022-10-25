package main

import "fmt"

func main() {

	//var conferenceName string = "Go Conference" equal conferenceName := "Go Conference"
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	// %v replace for varable
	// %T replace for type

	fmt.Printf("conferenceTickets is %T, remaining is %T, coonferenceName is %T\n", conferenceTickets, conferenceName, remainingTickets)

	fmt.Printf("welecom to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and  %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	//var name = "Go Conference"
	//Println = Print +\n
	//Printf  format for varable

	var userName string
	var userTickets int

	// ask user for their name

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
