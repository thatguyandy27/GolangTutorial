package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}

	fmt.Println("-------------")
	for i := 0; i < 1000; i++ {
		fmt.Println(i)
		if i == 5 {
			break
		}
	}
	fmt.Println("-------------")
	for i := 0; i < 3; i++ {
		if i == 1 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("-------------")

	i := 0
	for i < 3 {
		fmt.Println(i)
		i++

	}
}
