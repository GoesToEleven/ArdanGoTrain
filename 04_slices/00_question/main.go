// Declare a nil slice of integers. Declare a nil slice of integers. Create a
// loop that appends 10 values to the slice. Iterate over the slice and display
// each value. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main
import "fmt"

// Add imports.

// main is the entry point for the application.
func main() {
	// Declare a nil slice of integers.
	var numbers []int

	// Appends numbers to the slice.
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
	}

	// Display each value in the slice.
	for i, v := range numbers {
		fmt.Println(i, v)
	}

	// Declare a slice of strings and populate the slice with names.
	names := []string{"Jess", "Oliver", "Lillian", "Henry", "Padget"}

	// Display each index position and slice value.
	for i, v := range names {
		fmt.Println(i, v)
	}

	// // Take a slice of index 1 and 2 of the slice of strings.
	abridged := names[1:3]

	// Display each index position and slice values for the new slice.
	for i, v := range abridged {
		fmt.Println(i, v)
	}
}
