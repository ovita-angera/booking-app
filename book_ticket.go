package main

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
