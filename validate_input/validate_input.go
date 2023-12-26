package validateinput

import (
	"strings"
)

/*
This function is built specifically to check the validity of input data with the following criteria:
1. If either the first or last name are less than 2 letters, then the name is invalid and declared too short
2. If the email address does not contain an @ sign, then it is an invalid email address
3. If the number of tickets booked by the user are more than the remaining tickets, then the userTickets entered
is invalid
*/
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName bool
	var isValidEmail bool
	var isValidTickets bool

	if validateName(firstName) && validateName(lastName) {
		isValidName = true
	}

	if strings.Contains(email, "@") {
		isValidEmail = true
	}

	if userTickets <= remainingTickets {
		isValidTickets = true
	}

	return isValidName, isValidEmail, isValidTickets
}

func validateName(name string) bool {
	size := len(name) >= 2
	number := !strings.ContainsAny(name, "0123456789")

	return size && number
}
