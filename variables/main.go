package main

import (
	"fmt"
)

func main() {
	// Defining Variables
	var greet = "Hello world" // Throws error if not used. Variable type is inferred
	var greetAgain string = "Hello world again"

	// Defining multiple variables
	var (
		name string = "Daniel"
		age  int    = 23
	)

	// Shorthand Variable declaration
	school, matricNumber := "Uniuyo", "19/EG/ME/1361"

	fmt.Println(greet, greetAgain, name, age, school, matricNumber)

	const PUBLISHER = "publisher"
	fmt.Println(PUBLISHER)

	// Read data from console
	var read string
	fmt.Scan(&read)
	fmt.Println(read)
}
