package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(values)
	doubleAt(values, 1)
	fmt.Println(values)
	i := 10
	fmt.Println(i)
	double(i)
	fmt.Println(i)
	doublePtr(&i)
	fmt.Println(i)
}

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func doublePtr(i *int) {
	*i *= 2
}

func double(i int) {
	i *= 2
}
