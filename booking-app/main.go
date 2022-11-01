package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

// var conferenceName string = "Go Conference" equal conferenceName := "Go Conference" but global can't use :=
const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

// list of map
var bookings = make([]map[string]string, 0)

//var bookings = []string{}

func main() {

	greetUsers()

	for {

		userName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(userName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, userName, lastName, email)

			// call printFirstNames
			firstNames := getFirstNames()
			fmt.Printf("The first name of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {

			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("email address you entered is doesn't contain @ sign")
			}

			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is too short")
			}

			//fmt.Println("Your input data is invalid,try again")
			//fmt.Printf("We only have %v tickets , so you don't book %v tickets\n", remainingTickets, userTickets)
			//break
			//continue

		}
	}
}

func greetUsers() {
	fmt.Printf("Welecome to %v booking application\n", conferenceName)
	//fmt.Printf("welecom to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and  %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		//var names = strings.Fields(booking)
		//string.Fields(string) split the space
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
	//fmt.Printf("The first name of bookings are: %v\n", firstNames)

}

func getUserInput() (string, string, string, uint) {

	//bookings[0] = "Nana"
	var userName string
	var lastName string
	var email string
	var userTickets uint

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

	return userName, lastName, email, userTickets

}

func bookTicket(userTickets uint, userName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for a user
	// var myslice []string
	// var mymap map [string]string
	var userData = make(map[string]string)
	userData["userName"] = userName
	userData["lastName"] = lastName
	userData["email"] = email

	// we need convert the uint into a string
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	//bookings[0] = userName + " " + lastName
	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", userName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

}
