package main

import (
	"fmt"
)

func main() {
	// var x int
	// var x float64
	// var y float64
	// x = 1
	// y = 2
	x, y := 1., 2.
	// y := 2.
	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	// var mean float64
	// mean = (x + y) / 2
	mean := (x + y) / 2
	fmt.Printf("result: %v, type of %T\n", mean, mean)
}
