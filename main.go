package main

import (
	"fmt"
	"time"
)

// global variables
const conferenceTickets int = 50

type UserData struct {
	firstName    string
	lastName     string
	email        string
	ticketNumber uint
}

var conferenceName string = "Go conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

func main() {

	greetuser()

	for {
		//user input
		firstName, lastName, email, userTickets := getUserInput()
		//input validation
		isNameValid, isEmailValid, isTicketNumberValid, isTicketAvailable := ValidateUserInput(firstName, lastName, email, userTickets)

		if isNameValid && isEmailValid && isTicketNumberValid && isTicketAvailable {
			//book ticket
			userData := bookTicket(firstName, lastName, userTickets, email)

			//send ticket
			//using "go" to create a thread and make the app faster
			go sendTicket(userData)

			//retrive firstnames
			firstNames := getFirstnames()

			fmt.Printf("%v remaining tickets for %v\n", remainingTickets, conferenceName)
			fmt.Printf("These are all bookings %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("%v sold out! No tickets available!", conferenceName)
				break
			}

		} else {
			if !isNameValid {
				fmt.Println("Please, first name and last name should have more the 2 letters.")
			}

			if !isEmailValid {
				fmt.Println("Please, type a valid email address.")
			}

			if !isTicketNumberValid {
				fmt.Println("Please, number of tickets must be greater than zero.")
			}

			if !isTicketAvailable {
				fmt.Printf("There are only %v remaining tickets. You can buy %v tickets\n", remainingTickets, userTickets)
			}
		}
	}
}

func greetuser() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend.")
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Type your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Type your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Type your e-mail:")
	fmt.Scan(&email)
	fmt.Println("Type number of tickets to buy:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, userTickets uint, email string) UserData {
	remainingTickets = remainingTickets - userTickets

	//create user data map
	var userData = UserData{
		firstName:    firstName,
		lastName:     lastName,
		email:        email,
		ticketNumber: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is: %v\n", bookings)

	fmt.Printf("User %v %v booked %v tickets. You will receive a confirmation e-mail at %v\n", firstName, lastName, userTickets, email)
	return userData
}

func getFirstnames() []string {
	var firstnames []string
	for _, booking := range bookings {
		firstnames = append(firstnames, booking.firstName)
	}

	return firstnames
}

func sendTicket(userData UserData) {
	// simulate delay of the ticket generation
	time.Sleep(10 * time.Second)
	//send ticket
	fmt.Println("-------------------------")
	var ticket = fmt.Sprintf("%v tickets for costumer: %v %v ", userData.ticketNumber, userData.firstName, userData.lastName)
	fmt.Printf("Sending %v\n to e-mail %v\n", ticket, userData.email)
	fmt.Println("-------------------------")
}
