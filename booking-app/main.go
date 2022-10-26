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
	var lastName string
	var email string
	var userTickets int

	// ask user for their name

	// input output function and pointer
	fmt.Println("Enter your first name:")
	fmt.Scan(&userName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickts:")
	fmt.Scan(&userTickets)

	//fmt.Println(remainingTickets)
	//fmt.Println(&remainingTickets)

	//userName = "Tom"
	//userTickets = 2
	//fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", userName, lastName, userTickets, email)

}
