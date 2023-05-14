package main

import (
	"fmt"
)

func main() {

	// Go has primitive data types
	// uint8 - unsigned 1byte, uint16 - unsigned 2bytes, uint32 - unsigned 4bytes, uint64 - unsigned 8bytes, int8 - 1bytes, int16 - 2bytes, int32 - 4bytes, int64 - 8bytes
	// float32 - single precision, float64 - double precision, complex64, complex128

	var age uint = 459999999
	fmt.Println(age)

	var science float32 = 1234.89
	fmt.Println(science)

	// var name string = "Daniel"[1]
	fmt.Println("Daniel"[1])

	// Boolean
	fmt.Println(true && true)  // returns true
	fmt.Println(true && false) // returns false
	fmt.Println(true || true)  // returns true
	fmt.Println(true || false) // returns false
	fmt.Println(!true)         // returns false

}
