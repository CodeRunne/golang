package main

import (
	"fmt"
	"math"
)

// return multiple results
func swap(x, y string) (string, string) {
	return y, x
}

func add(x, y int) int {
	return (x + y)
}

// Go named return values, naked return statement(return statement without argument).
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// Function Values -- functions can be used as function arguments and return values
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Function closures - It's a function value that references variables from outside the function
func adder() func(int) int { //returns a function value
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// Fibonacci Closure
func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}

func main() {
	a, b := swap("hello", "world")
	c, d := split(5)
	fmt.Println(a, b, add(2, 4), c, d)

	//Anonymous functions
	getName := func(name string) string {
		return "My name is " + name
	}

	fmt.Println(getName("dan"))

	// Function Value
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Printf("%.2f\n", hypot(14, 3))
	fmt.Println(compute(hypot))    // Uses the value defined in the compute function - returns 5
	fmt.Println(compute(math.Pow)) // Uses the value defined in the compute function - returns 81

	pos, neg := adder(), adder()
	for i := 0; i <= 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}

}
