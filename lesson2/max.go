package main

import "fmt"

func main() {
	nums := []int{16, 5, 32, 234, 0, 123, 43}
	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	fmt.Println(max)
}
