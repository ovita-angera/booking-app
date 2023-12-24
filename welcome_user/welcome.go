package welcomeuser

import "fmt"

// Function to greet the user
func GreetUser(name string) {
	fmt.Printf("Welcome to the %v ticket booking application.\n", name)
}

// This function prints out information about the total and remaining number of tickets
func PrintInfo(conf_tickets uint, rem_tickets uint) {
	fmt.Printf("There is a total of %v tickets, and %v are still available.\n", conf_tickets, rem_tickets)
	fmt.Println("Get your tickets here to attend")
}
