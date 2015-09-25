// Write a program where two goroutines pass an integer back and forth
// ten times. Display when each goroutine receives the integer. Increment
// the integer with each pass. Once the integer equals ten, terminate
// the program cleanly.
package main

import (
	"fmt"
	"sync"
)

// Declare a wait group variable.
var wg sync.WaitGroup

// main is the entry point for all Go programs.
func main() {
	// Declare and initialize an unbuffered channel
	// of type int.
	share := make(chan int)

	// Increment the wait group for each goroutine
	// to be created.
	wg.Add(2)

	// Create the two goroutines.
	go goroutine("Todd", share)
	go goroutine("Jess", share)

	// Send the initial integer value into the channel.
	share <- 1

	// Wait for all the goroutines to finish.
	wg.Wait()
}

// Declare a function for the goroutine. Pass in a name for the
// routine and a channel of type int used to share the value back and forth.
func goroutine(name string, share chan int) {
	for {
		// Receive on the channel, waiting for the integer.
		value, ok := <-share

		// Check if the channel is closed.
		if !ok {
			// Call done on the waitgroup.
			wg.Done()
			// Display the goroutine is finished and return.
			fmt.Println("Go routine is down.")
			return
		}

		// Display the integer value received.
		fmt.Printf("Goroutine %s Inc %d\n", name, value)
		// Check is the value from the channel is 10.
		if value == 10 {
			// Close the channel.
			close(share)
			// Call done on the waitgroup.
			wg.Done()
			// Display the goroutine is finished and return.
			fmt.Printf("Goroutine %s finished\n", name)
			return
		}

		// Increment the value by one and send in back through
		// the channel.
		share <- (value + 1)
	}
}
