package main

// import the required packages for the booking logic
import (
	// a module within our program
	getinput "booking-app/get_input"
	sendticket "booking-app/send_ticket"
	validateinput "booking-app/validate_input"
	welcomeuser "booking-app/welcome_user"
	"fmt" // a module from Go packages called the `format` package
	// a module from Go
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

type UserData struct {
	FirstName       string
	LastName        string
	Email           string
	NumberOfTickets uint
}

var conferenceTickets uint = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

/*
Define a data structure which can hold values of different types
In this case, we create a user data structure to store user information when booking a ticket
*/

// The entry point for execution of our booking app CLI app
func main() {

	welcomeuser.GreetUser(conferenceName)

	// Ticket booking logic executed in a loop
	for {
		welcomeuser.PrintInfo(conferenceTickets, remainingTickets)

		firstName, lastName, email, userTickets := getinput.GetUserInput()
		isValidName, isValidEmail, isValidTickets := validateinput.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTickets {
			BookTicket(firstName, lastName, email, userTickets)
			sendticket.SendTicket(firstName, lastName, email, userTickets)

			firstNames := GetFirstNames()
			fmt.Printf("Bookings: %v\n", bookings)
			fmt.Printf("The first names of the bookings are: %v\n", firstNames)

			remainingTickets = remainingTickets - userTickets

			// condition to end the program
			if remainingTickets == 0 {
				fmt.Printf("The %v conference is totally booked out, try again next year.\n", conferenceName)
				break
			}

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
	}
}

// This function extracts the first names of the bookings and returns a list of the first names
func GetFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.FirstName)
	}
	return firstNames
}

// This function executes the booking ticket logic
func BookTicket(firstName string, lastName string, email string, userTickets uint) {
	bookings = append(bookings, UserData{
		FirstName:       firstName,
		LastName:        lastName,
		Email:           email,
		NumberOfTickets: userTickets,
	})
}
