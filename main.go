package main

import (
	"fmt"
	"sync"
	"time"
)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// sync package to achieve syunchronization
var wg = sync.WaitGroup{}

func main() {
	var conferenceName = "Go conference"
	const conferenceTicket = 50
	var remainingTicket uint = 50
	bookings := make([]UserData, 0)

	fmt.Printf("Welcome to our %s booking application \n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still available \n", conferenceTicket, remainingTicket)
	fmt.Println("Get your tickets here")

	for {

		if remainingTicket == 0 {
			fmt.Println("Our conference is sold out.")
			break
		}

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNo := validateUserInput(firstName, lastName, email, userTickets, remainingTicket)
		if isValidName && isValidEmail && isValidTicketNo {

			remainingTicket = remainingTicket - userTickets

			//create map for users
			var userData = UserData{
				firstName:       firstName,
				lastName:        lastName,
				email:           email,
				numberOfTickets: userTickets,
			}
			//bookings list of key value pairs
			bookings = append(bookings, userData)

			fmt.Printf("Thank you %v %v for booking %v tickets, you will recieve confirmation in your %v \n",
				firstName, lastName, userTickets, email)

			//go keyword for concurrency starting new thread or goroutine
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			fmt.Printf("Remianing tickets for conference %v \n", remainingTicket)

			fmt.Printf("All our bookings %v", bookings)

		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email you entered is invalid")
			}
			if !isValidTicketNo {
				fmt.Println("Number of tickets is invaid")
			}
			fmt.Println("Try continuing again")

		}
		//
		wg.Wait()
	}

}
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scanln(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scanln(&lastName)
	fmt.Println("Enter your email address:")
	fmt.Scanln(&email)
	fmt.Println("Enter number of tickets")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	//Simulating interruption or wait time
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#########")
	fmt.Printf("Sending ticket:\n %v to email address %v", ticket, email)
	fmt.Println("#########")
	//
	wg.Done()
}
