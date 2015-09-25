// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

import "fmt"

// Add imports.

// Declare a type named user.
type user struct {
	name  string
	email string
	age   int
}

// Create a function that changes the value of one of the user fields.
/* add pointer parameter, add value parameter */
func update(u *user, newAge int) {
	u.age = newAge
	// Use the pointer to change the value that the
	// pointer points to.
}

// main is the entry point for the application.
func main() {
	// Create a variable of type user and initialize each field.
	bill := user{
		name:  "Bill",
		email: "bill@ardanstudio.com",
		age:   38,
	}

	// Display the value of the variable.
	fmt.Println(bill.age)

	// Share the variable with the function you declared below.
	update(&bill, 39)

	// Display the value of the variable.
	fmt.Println(bill.age)
}
