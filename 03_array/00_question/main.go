// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
package main
import "fmt"

// Add imports.

// main is the entry point for the application.
func main() {
	// Declare an array of 5 strings set to its zero value.
	var names [5]string
	// Declare an array of 5 strings and pre-populate it with names.
	friends := [5]string{"Jess", "Oliver", "Lillian", "Youness", "Greg"}
	// Assign the populated array to the array of zero values.
	names = friends
	// Iterate over the first array declared.
	for i, friend := range names {
		fmt.Println(i, " - ", friend)
	}
}