// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a players batting average. The formula is Hits / AtBats.
// Declare a slice of this type and initalize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

import "fmt"

// Add imports.

// Declare a struct that represents a ball player.
// Include field called name, atBats and hits.
type player struct {
	name   string
	atBats int
	hits   int
}

// Declare a method that calculates the batting average for a batter.
func (p player) average() float64 {
	return (float64(p.hits) / float64(p.atBats)) * 1000
}

// main is the entry point for the application.
func main() {
	// Create a slice of players and populate each player
	// with field values.
	players := []player{
		{"Sandy", 22, 18},
		{"Babe", 44, 40},
		{"Reggie", 38, 18},
	}

	// Display the batting average for each player in the slice.
	for _, v := range players {
		fmt.Println(v.name, " - ", v.average())
	}
}
