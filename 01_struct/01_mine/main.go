// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/ItPe2EEy9X

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

import "fmt"

// Add imports.

// Add user type and provide comment.
type usr struct {
	name  string
	email string
	age   int
}

// main is the entry point for the application.
func main() {

	// Declare variable of type user and init using a struct literal.
	bill := usr{
		"Bill",
		"Bill@ardanstudios.com",
		38,
	}
	// Display the field values.
	fmt.Println(bill.name)
	fmt.Println(bill.email)
	fmt.Println(bill.age)
	// Declare a variable using an anonymous struct.
	ed := struct {
		name  string
		email string
		age   int
	}{
		"ed",
		"ed@ardanstudios.com",
		42,
	}

	// Display the field values.
	fmt.Println(ed.name)
	fmt.Println(ed.email)
	fmt.Println(ed.age)
}
