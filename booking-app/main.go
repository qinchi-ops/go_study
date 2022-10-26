package main

import (
	"fmt"
	"strings"
)

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

	// for ture = for
	// for remainingTickets > 0 && len(bookings) < 50 {}
	for {

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

		if userTickets <= remainingTickets {
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

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				//string.Fields(string) split the space
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first name of bookings are: %v\n", firstNames)

			//var noTicketsRemaining bool = remainingTickets == 0
			//noTicketsRemaining := remainingTickets == 0
			//if noTicketsRemaining {
			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {

			fmt.Printf("We only have %v tickets , so you don't book %v tickets\n", remainingTickets, userTickets)
			//break
			continue

		}

	}

}
