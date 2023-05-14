package main

import (
	"fmt"
)

type Person struct {
	name           string
	age            int
	jobDescription string
}

type Vertex struct {
	x, y int
}

func main() {

	// Defining Structs
	// structs (Structure) is used for storing a collection of members of different datatype into a single variable
	// syntax => type struct_name struct {
	// 	member1 datatype1
	// 	member2 datatype2
	// 	member3 datatype3
	// }

	var pers1 Person
	pers1.name = "Daniel Ibok"
	pers1.age = 23
	pers1.jobDescription = "Software Developer"

	fmt.Printf("%v\t%T\n", pers1, pers1)

	// Passing a struct to a function
	get(pers1)

	fmt.Println(Vertex{1, 4})

	v := Vertex{4, 5}
	fmt.Println(v.x)

	// pointer to structs
	c := Vertex{5, 8}
	p := &c
	p.x = 1e9
	fmt.Println(p.x) // prints 5

	// Struct Literals
	// A struct literal denotes a newly allocated struct values of its fields.
	v1 := Vertex{1, 2}  // has type vertex
	v2 := Vertex{x: 1}  // y:0 is implicit
	v3 := Vertex{}      // x:0 and y:0
	p1 := &Vertex{1, 2} //has type vertex
	fmt.Printf("%v\t%v\t%v\t%v\n", v1, v2, v3, p1)
}

// method
func get(pers Person) {
	fmt.Printf("Name: %v\n", pers.name)
	fmt.Printf("Age: %v\n", pers.age)
	fmt.Printf("Job Description: %v\n", pers.jobDescription)
}
