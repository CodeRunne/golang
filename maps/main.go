package main

import (
	"fmt"
)

func main() {
	// Defining Maps
	// --- Maps are used for storing data values in key:value pairs. The default value for a map is nil
	//
	// ---- METHOD 1 ---------
	// syntax => map[keyDataType]valueDataType {key1:value1, key2:value2}

	mymap := map[string]string{"Name": "John Doe", "Age": "32", "Job": "Mechanical Engineer"}

	// Modifying the value of an element
	mymap["Age"] = "45"        // changes "Age" to 45
	mymap["crypt"] = "bitcoin" // adds a new key:value pair to the map

	fmt.Println(mymap, len(mymap))

	// ---- METHOD 2 -----------
	// syntax => make(map[keyDataType]valueDataType)

	mymap2 := make(map[string]int) // Creating an empty map using the make function

	// Populating the map
	mymap2["style"] = 2
	mymap2["groove"] = 3

	fmt.Printf("mymap2\t%v\n", mymap2)

	// ---- METHOD 3 -----------
	// syntax => var a map[keyDataType]valueDataType -- the assignment operator is omitted

	var mymap3 map[string]string // An empty map
	fmt.Println(mymap3)

	// CHECK TO CONFIRM IF mymap2 and mymap3 is nil
	fmt.Println(mymap2 == nil)
	fmt.Println(mymap3 == nil)

	// Modify mymap3
	//mymap3["name"] = "dan" // throws an error, can't assign to a nil map
	//fmt.Println(mymap3)

	// ALLOWED KEY TYPES
	// 1 - ARRAYS
	// 2 - BOOLEANS
	// 3 - NUMBERS
	// 4 - POINTERS
	// 5 - STRUCTS
	// 6 - INTERFACES
	// 7 - STRINGS

	// ALLOWED VALUE TYPES
	// ACCEPTS ALL DATA TYPES

	// ---- DELETE ELEMENTS IN A MAP
	// syntax => delete(map_name, key)

	mymap4 := map[string]string{"age": "45", "name": "dan", "job": "Mechanical Engineer", "experience": ""}
	fmt.Println(mymap4, len(mymap4))

	delete(mymap4, "age") // deletes the age key value pair using the key
	fmt.Println(mymap4, len(mymap4))

	// CHECK FOR SPECIFIC ELEMENTS IN THE MAP
	val1, ok1 := mymap4["name"]   // stores value in the val1 variable and true in the ok1 variable if the value of the key exists
	val2, ok2 := mymap4["colors"] // check for non-existant key
	val3, ok3 := mymap4["experience"]
	_, ok4 := mymap4["job"] // only checking for existing key

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)

	// when checking for a key existence always check if it true using the second value that'll be returned by go to preven bug
	if _, ok := mymap4["colors"]; ok { //shorthand method of checking
		fmt.Println("Value was found")
	} else {
		fmt.Println("Value not found")
	}
}
