package getinput

import "fmt"

// This function gets the user input and returns the input
func GetUserInput() (string, string, string, uint) {
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
