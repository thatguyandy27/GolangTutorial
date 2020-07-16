package main

import (
	"fmt"
)

func safeValue(vals []int, index int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("ERROR: %s\n", err)
		}
	}()
	return vals[index]
}

func main() {
	// vals := []int{1, 2, 3}
	// v := vals[10]
	// fmt.Println(v)
	// file, err := os.Open("FOOBAZ")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	vals := []int{1, 2, 3}
	v := safeValue(vals, 10)
	fmt.Println(v)
}
