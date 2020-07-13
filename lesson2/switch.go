package main

import (
	"fmt"
)

func main() {
	x := 2
	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}

	switch {
	case x > 100:
		fmt.Println("x is very big")
	case x > 10:
		fmt.Println("x is moderate")
	default:
		fmt.Println("whatev")
	}
}
