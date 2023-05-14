package main

import (
	"fmt"
	"strings"
)

func main() {
	// ------------- String functions --------------
	fmt.Println(strings.ToUpper("hello world")) // Convert to uppercase
	fmt.Println(strings.ToLower("HELLO WORLD")) // Convert to lowercase

	// Raw Strings
	mystring := `This life is full of challenges, but never give up!!!` // using back ticks
	fmt.Println(mystring)

	// concatenating strings
	fmt.Println(mystring + " and always believe in God")

	// Looping through a string
	utf := "Hello, 世界" //UTF characters

	for i, c := range utf {
		fmt.Printf("%d: %s\n", i, string(c))
	}

	fmt.Println(strings.HasPrefix(mystring, "T"))   // Checks if the string starts with the letter `T`
	fmt.Println(strings.HasSuffix(mystring, "!"))   // Checks if the string ends with the character `!`
	fmt.Println(strings.Contains(mystring, "full")) // Checks to see if the string contain the word
	fmt.Println(strings.Count(mystring, "f"))       // returns the number of occurence of the character or word in the string

	// ----------- STRING MANIPULATION -----------
	// -- strings.Join(), strings.Split(), strings.Fields(), strings.ReplaceAll()
	// 1 - strings.Join(slice, seperator) - explode or combines a slice of string to a single string

	s := strings.Join([]string{"Dan", "Joe", "Jen", "wen"}, ", ")
	fmt.Println(s) // returns Dan, Joe, Wen
	fmt.Printf("%v %T\n", s, s)

	// 2 - strings.split(string, delimiter) - implodes or breaks a string into a slice

	ballon := "I want my ballon" // the whitespace is the delimiter
	newSlice := strings.Split(ballon, " ")
	fmt.Println(newSlice) //returns [I want my balllon]
	fmt.Printf("%q %T\n", newSlice, newSlice)

	// 3 - strings.Fields(string) - Similar to string.Split() only that it ignores all whitespace and returns only the fields

	detail := "My name   is  John    Doe"
	fmt.Println(strings.Fields(detail)) // returns [My name is John Doe]
	fmt.Printf("%q\n", strings.Fields(detail))

	// 4 - strings.ReplaceAll(string, search, replace)
	fmt.Println(strings.ReplaceAll(ballon, "my", "your")) // returns I want your ballon
}
