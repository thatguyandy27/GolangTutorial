package main

import (
	"fmt"
	"strconv"
)

func main() {
	const max = 20

	for i := 1; i <= max; i++ {
		word := ""
		if i%3 == 0 {
			word = "fizz"
		}
		if i%5 == 0 {
			word = word + "buzz"
		}
		if word == "" {
			word = strconv.Itoa(i)
		}

		fmt.Println(word)
	}
}
