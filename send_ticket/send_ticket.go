package sendticket

import (
	"fmt"
)

// This function prints out a message sending the ticket to a particular user
func SendTicket(firstName string, lastName string, email string, userTickets uint) {
	ticket := fmt.Sprintf("%v tickets for %v %v sent to the email address %v", userTickets, firstName, lastName, email)
	fmt.Println("******************************************")
	fmt.Printf("Sending ticket...\n%v\n", ticket)
	fmt.Println("******************************************")
}
