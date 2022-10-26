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

	//var bookings = [50]string{"nana","Nicole","Peter"}

	//var bookings []string
	//also can use bookings :=[]string{}
	bookings := []string{}

	//bookings[0] = "Nana"
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

	remainingTickets = remainingTickets - userTickets
	//bookings[0] = userName + " " + lastName
	bookings = append(bookings, userName+" "+lastName)

	//fmt.Println(remainingTickets)
	//fmt.Println(&remainingTickets)

	//userName = "Tom"
	//userTickets = 2
	//fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

	fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("slice type: %T\n", bookings)
	fmt.Printf("slice length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", userName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

	fmt.Printf("These are all our bookings:\n", bookings)
}
