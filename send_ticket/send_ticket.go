package sendticket

import (
	"fmt"
	"time"
)

// This function prints out a message sending the ticket to a particular user
func SendTicket(firstName string, lastName string, email string, userTickets uint) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v to the email address %v", userTickets, firstName, lastName, email)
	fmt.Println("******************************************")
	fmt.Printf("Sending ticket...\n%v\n", ticket)
	fmt.Println("******************************************")
}
