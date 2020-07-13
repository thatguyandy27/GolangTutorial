package main

import (
	"fmt"
)

func main() {
	x := 10

	if x > 5 {
		fmt.Println("x is big")
	}

	if x > 100 {
		fmt.Println("X is very big")
	} else {
		fmt.Println("x is not very big")
	}

	if x > 5 && x < 15 {
		fmt.Println("x is just right")
	}

	if x < 20 || x > 30 {
		fmt.Println("x is out of range")
	}

	a := 11.
	b := 20.

	if frac := a / b; frac > .5 {
		fmt.Println("a is > .5 of b")
	}
}
