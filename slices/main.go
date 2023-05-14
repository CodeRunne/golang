package main

import (
	"fmt"
	"strings"
)

func main() {
	// Slices in Golang
	// Slice is similar to array, but is more powerful and flexible -- The length of a slice can grow and shrink
	//There are three(3) ways of defining a slice

	// ---- METHOD 1 ------
	//Syntax => []datatype{values}

	myslice := []int{1, 2, 3, 4}
	fmt.Println(myslice, len(myslice))
	fmt.Printf("%T\n", myslice)

	//---- METHOD 2 -------
	// syntax => myarr[start:end] - creating a slice from an array
	myarr := [3]int{1, 2, 3}
	myslice2 := myarr[1:2]
	fmt.Println(myslice2)

	// ---- METHOD 2 -------
	// syntax => make([]datatype, length, capacity(optional))
	myslice3 := make([]int, 5)
	myslice3[2] = 4
	fmt.Println(myslice3)

	var myslice4 []int
	fmt.Println(myslice4)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	var t map[string]int
	fmt.Printf("value: %v,\ttype:%T\n", t, t)

	h := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Printf("value: %v,\ttype: %T\n", h, h)

	u := []int{2, 3, 5, 7, 11, 13}

	u = u[:]
	fmt.Printf("%v\n", u)

	u = u[:3]
	fmt.Printf("%v\n", u)

	u = u[1:]
	fmt.Printf("%v\n", u)

	u = u[:0]
	fmt.Printf("%v\n", u)

	u = u[:4]
	fmt.Printf("%v\n", u)

	// Slices of slices
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var appends []int
	appends = append(appends, 2)
	appends = append(appends, 4, 6, 8)
	fmt.Println(appends)

}
