package main

import (
	"github.com/vikaspoddar/booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings []User

const conferenceTickets = 50

var wg = sync.WaitGroup{}

type User struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTickets(&remainingTickets, userTickets, &bookings, firstName, lastName, email, conferenceName)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)
		firstNames := getFirstName(bookings)
		fmt.Printf("The first name of bookings %v\n", firstNames)
		if remainingTickets == 0 {
			fmt.Printf("Our conference is booked out. Come next year.")

		}
	} else {

		if !isValidName {
			fmt.Println("Enter your name correctly")
		}
		if !isValidEmail {
			fmt.Println("Enter a valid email address")
		}
		if !isValidTicketNumber {
			fmt.Println("Enter a valid number of tickets")
		}
		fmt.Printf("Your input data is invalid. Try again.\n")
	}
	wg.Wait()
}

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstName(bookings []User) []string {
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Printf("Enter your first name: ")
	_, _ = fmt.Scan(&firstName)
	fmt.Printf("Enter your last name: ")
	_, _ = fmt.Scan(&lastName)
	fmt.Printf("Enter your email address: ")
	_, _ = fmt.Scan(&email)
	fmt.Printf("Enter number of tickets: ")
	_, _ = fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTickets(remainingTickets *uint, userTickets uint, bookings *[]User, firstName, lastName, email, conferenceName string) {
	*remainingTickets = *remainingTickets - userTickets
	var userData = User{firstName, lastName, email, userTickets}
	*bookings = append(*bookings, userData)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", *remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName, lastName, email string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v Ticket for %v %v", userTickets, firstName, lastName)
	fmt.Println("###########")
	fmt.Printf("Sending %v \n to email address %v\n", ticket, email)
	fmt.Println("###########")
	wg.Done()
}
