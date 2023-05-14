package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// Casting in go
	var num uint8 = 45
	float := float32(num)
	float1 := float64(float)
	num1 := int(float1) // uint8 - 64 | int8 - 64

	fmt.Printf("value: %d, type: %T\n", num, num)
	fmt.Printf("value: %.2f, type: %T\n", float, float)
	fmt.Printf("value: %.2f, type: %T\n", float1, float1)
	fmt.Printf("value: %d, type: %T\n", num1, num1)

	// String Conversion
	// converting numbers to string using the strconv.Itoa() -- strconv package
	num2 := 34
	stringConv := strconv.Itoa(num2)
	fmt.Printf("value: %#v, type: %T\n", stringConv, stringConv)

	// Converting string to number using the strconv.Atoi()
	string1 := "50"
	conv, err := strconv.Atoi(string1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("value: %#v, type: %T\n", conv, conv)

	// -----------------------------------------------

	// converting string to float using the strconv.ParseFloat()
	// float2 := "45.6"
	// conv1, err1 := strconv.ParseFloat(float2)

	// if err1 != nil {
	// 	log.Fatal(err1)
	// }

	// fmt.Printf("value: %#v, type: %T\n", conv1, conv1)

	// ------------------------------------------------

	// converting string to byte -- in go strings are stored as a slice of bytes

	a := "hello there, young man"
	b := []byte(a) // convert to byte
	c := string(b) // convert back to string

	fmt.Printf("%v\n%v\n%v\n", a, b, c)

	// converting float to string using fmt.Sprint(float)
	float3 := 34.5
	fmt.Printf("value: %#v, type: %T\n", fmt.Sprint(float3), fmt.Sprint(float3))
}
