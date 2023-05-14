package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Go Methods - It's just a function with a special receiver argument.
// The receiver appears in it's own argument list btw the func keyword and the method name.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// A method can also be written as a regular function
func Raise(p Vertex) float64 {
	return math.Pow(p.X, p.Y)
}

// A method can also be a non-struct type
type MyFloat float64

func (f MyFloat) Float() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	p := Vertex{13, 2}
	fmt.Println(Raise(p))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Float())
}
