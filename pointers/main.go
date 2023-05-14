package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// pointers and functions
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Scaler(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScalerFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	// GO has no pointer arithmetic
	// Pointer hold the memory address of a value
	// & - is for getting the address of a value
	// * - for dereferencing a pointer and assigning a pointer

	var p *int
	i := 43
	p = &i

	fmt.Println(p, *p)

	v := Vertex{3, 4}
	Scale(&v, 20)
	fmt.Println(v.Abs())

	v1 := Vertex{6, 5}
	v1.Scaler(4)
	ScalerFunc(&v1, 10)
	fmt.Println(v1)

	v2 := &Vertex{7, 3}
	v2.Scaler(4)
	ScalerFunc(v2, 10)
	fmt.Println(v2)
}
