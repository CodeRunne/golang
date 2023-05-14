package main

import (
	"fmt"
)

func main() {
	// Defer postpones a line of code or function until all surrounding function are executed
	defer fmt.Println("hello") // postpones the function
	fmt.Println("world")

	// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}
