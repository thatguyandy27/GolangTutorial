package main

import "fmt"

func cleanup(a string) {
	fmt.Printf("Cleaning up %s\n", a)
}
func worker() {
	defer cleanup("A")
	fmt.Println("worker")

}

func main() {
	worker()
}
