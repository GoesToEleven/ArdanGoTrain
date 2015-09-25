// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
package main

import "fmt"

// main is the entry point for the application.
func main() {
	// Declare string arrays to hold names.
	var names [5]string

	// Declare an array pre-populated with friend's names.
	friends := [5]string{"Joe", "Ed", "Jim", "Erick", "Bill"}

	// Asssign the array of friends to the names array.
	names = friends

	// Display each name in names.
	for i, name := range names {
		fmt.Println(name, &name, &names[i])
	}

	//	Joe 	0x82024e220 0x82025c0a0
	//	Ed 		0x82024e220 0x82025c0b0
	//	Jim 	0x82024e220 0x82025c0c0
	//	Erick 	0x82024e220 0x82025c0d0
	//	Bill 	0x82024e220 0x82025c0e0

}
