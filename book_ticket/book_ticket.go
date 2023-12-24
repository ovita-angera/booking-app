package bookticket

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// This function extracts the first names of the bookings and returns a list of the first names
func GetFirstNames(bookings []UserData) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// This function executes the booking ticket logic
func BookTicket(bookings []UserData, firstName string, lastName string, email string, userTickets uint) {
	bookings = append(bookings, UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	})
}
