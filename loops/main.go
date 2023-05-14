package main

import (
	"fmt"
)

func main() {
	// Go uses only the for loop
	// THREE WAYS OF USING THE FOR LOOP
	// for statement1; statement2; statement3 { }
	// for key, val := range arr { }
	// for statement { }

	// ----- METHOD 1 -----------
	for i := 0; i <= 9; i++ {

		if i == 4 {
			continue
		}

		if i == 7 {
			break
		}

		fmt.Println(i)
	}

	// ----- METHOD 2 ------------
	// syntax => for statement { } - similar to while loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// ----- METHOD 3 ------------
	// Specifically for looping through array|slice|maps

	myarr := [4]string{"dan", "joe", "henry", "renold"}
	for idx, val := range myarr {
		fmt.Println(idx, val)
	}

	// omitting the array key
	for _, val := range myarr {
		fmt.Println(val)
	}

	myslice := []int{100, 200, 300, 400, 500}
	fmt.Println(myslice)
	for id, val := range myslice {
		fmt.Println(id, val)
	}

	// Modifying using for loop
	for i := 0; i < len(myslice); i++ {
		myslice[i] = myslice[i] * 2
		fmt.Printf("\n%v", myslice[i])
	}

	mymap := map[string]string{"name": "joe", "age": "34", "job": "developer"}
	for id, val := range mymap {
		fmt.Println(id, val)
	}

	// Forever Loop
	count := 0
	for {
		fmt.Println(count)
		if count == 100 {
			break
		}
		count++
	}
}
