# booking-app
This is a command line application that simulates a ticket booking application.
It has great inspiration from the youtube course Golang Tutorial for Beginners | Full Go Course (https://www.youtube.com/watch?v=yyUHQIec83I). 
The code implements several golang coding principles and even extends to indicate the functionality of several modules in the same application. It is a good blue print to start from in order to undestand and even grow a sample UI in time (a work in progress). 
The application logic is pretty straightforward:
- First to run the app, once in the app folder using the terminal run `go run .` which consolidates the logic and harmony of all the files in the directory and runs the app as a whole. In this case, the `main.go` file borrows input validation logic from `helper.go` in the helper directory.
- Once in the app, the logic is:
  1. Greet the user
  2. Invoke a loop that
     a. Prompts the user for input - first name, last name, email and number of tickets they wish to book
     b. Validates the input using logic explained in the `helper.go` file
     c. Then the app updates a list of bookings with the information
     d. The app checks if there are tickets still available for booking
     - if there are any, the loop continue
     - else, the loop breaks and the app quits
     e. The app prints out information on the bookings, such as the first names and number of remaining tickets

The app implements use of `structs`, in this case `userData`.
