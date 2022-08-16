package main

import "strings"

func ValidateUserInput(firstname string, lastname string, email string, userTickets uint) (bool, bool, bool, bool) {
	isNameValid := len(firstname) > 2 && len(lastname) > 2
	isEmailValid := strings.Contains(email, "@")
	isTicketNumberValid := userTickets > 0
	isTicketAvailable := userTickets < remainingTickets

	return isNameValid, isEmailValid, isTicketNumberValid, isTicketAvailable
}
