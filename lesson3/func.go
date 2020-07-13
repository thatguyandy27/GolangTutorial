package main

import "fmt"

func main() {
	val := add(1, 2)
	fmt.Println(val)

	d, r := divmod(5, 2)

	fmt.Println(d)
	fmt.Println(r)
}

func add(a int, b int) int {
	return a + b
}

func divmod(a int, b int) (int, int) {
	return a / b, a % b
}
