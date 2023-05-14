package main

import (
	"fmt"
)

func main() {
	fmt.Println("Please enter your name: ")
	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		panic(err)
	}
}
