package main

// import the required packages for the booking logic
import (
	"booking-app/helper" // a module within our program
	"fmt"                // a module from Go packages called the `format` package
	"time"               // a module from Go
)

// declare the package level variables
/*
This makes sure all the functions under this main package have access to these variables
> conferenceName - Name of the conference with available tickets, in this case: Go Conference
> conferenceTickets - The total number of tickets available for the said conference
> remainingTickets - This is a value place holder that would update each time a user books tickets
> bookings - this is a slice to keep a record of each booking's information
*/
const conferenceName string = "Go Conference"

var conferenceTickets uint = 50
var remainingTickets uint = 50
var bookings = make([]userData, 0)

/*
Define a data structure which can hold values of different types
In this case, we create a user data structure to store user information when booking a ticket
*/
type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// The entry point for execution of our booking app CLI app
func main() {

	greetUser()

	// Ticket booking logic executed in a loop
	for {
		printInfo()

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTickets {
			bookTicket(firstName, lastName, email, userTickets)
			sendTicket(firstName, lastName, email, userTickets)

			firstNames := getFirstNames()
			fmt.Printf("The first names of the bookings are: %v\n", firstNames)

			// condition to end the program
			if remainingTickets == 0 {
				fmt.Printf("The %v conference is totally booked out, try again next year.\n", conferenceName)
				break
			}

			remainingTickets = remainingTickets - userTickets

		} else {
			if !isValidName {
				fmt.Println("The name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("The email you entered does not contain an @")
			}

			if !isValidTickets {
				fmt.Printf("There are only %v tickets available, you booked %v tickets.\n", remainingTickets, userTickets)
			}
		}
		fmt.Printf("Bookings: %v\n", bookings)
	}
}

// Function to greet the user
func greetUser() {
	fmt.Printf("Welcome to the %v ticket booking application.\n", conferenceName)
}

// This function prints out information about the total and remaining number of tickets
func printInfo() {
	fmt.Printf("There is a total of %v tickets, and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// This function extracts the first names of the bookings and returns a list of the first names
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// This function gets the user input and returns the input
func getUserInput() (string, string, string, uint) {
	// declare variables
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// prompt user for input and save the input respectively
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you wish to book: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// This function executes the booking ticket logic
func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	bookings = append(bookings, userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	})
}

// This function prints out a message sending the ticket to a particular user
func sendTicket(firstName string, lastName string, email string, userTickets uint) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v to the email address %v", userTickets, firstName, lastName, email)
	fmt.Println("******************************************")
	fmt.Printf("Sending ticket...\n%v\n", ticket)
	fmt.Println("******************************************")
}
