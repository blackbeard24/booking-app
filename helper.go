package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTicket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNo := userTickets > 0 && userTickets <= remainingTicket
	return isValidName, isValidEmail, isValidTicketNo
}
