package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50

// var bookings = make([]map[string]UserData, 0)
var bookings = []UserData{}

var wg = sync.WaitGroup{}

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

func main() {

	greetUsers()

	// var bookings [50]string // Array
	for {
		// Slice
		//var bookingIndex int = 0
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			fmt.Println("Thank you for your input!")
		} else { // go syntax rule that else be in same line as if closing bracket
			if !isValidName {
				fmt.Println("First name or last name is too short!")
			}
			if !isValidEmail {
				fmt.Println("Email address is not valid!")
			}
			if !isValidTicketNumber {
				fmt.Printf("Number of tickets you want to book should be between 1 and %v\n", remainingTickets)
			}
			continue // skip the rest of the loop and start over
		}
		bookTickets(userTickets, firstName, lastName, email)
		//fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		firstNames := printFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
		if remainingTickets == 0 {
			fmt.Println("All tickets are sold out!")
			break
		}

	}
	wg.Wait() // Wait for all goroutines to finish
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Grab your tickets here!")
}

func printFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		//var names = strings.Fields(booking.firstName) // basically string.split
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets you want to book:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets int, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - uint(userTickets)
	//bookings[bookingIndex] = firstName + " " + lastName
	//bookingIndex++
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	// bookings = append(bookings, userData)

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	wg.Add(1) // Increment the WaitGroup counter
	go sendTicket(userTickets, firstName, lastName, email)
	fmt.Printf("Print Whole Array: %v\n", bookings)
	// fmt.Printf("The first value: %v \n", bookings[0])
	// fmt.Printf("Array type: %T\n", bookings)
	// fmt.Printf("The length of the array: %v\n", len(bookings))
	fmt.Printf("Thank you  %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
}

func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second) // Simulating a delay for sending the ticket
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("----------------------------------")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Printf("----------------------------------\n")
	wg.Done() // Indicate that this goroutine is done
}
