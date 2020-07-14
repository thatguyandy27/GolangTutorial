package main

import (
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
}

func (p *Point) Move(dx int, dy int) {
	p.x += dx
	p.y += dy
}

type Square struct {
	center Point
	length int
}

func (s *Square) Move(dx int, dy int) {
	s.center.Move(dx, dy)
}

func (s *Square) Area() int {
	return s.length * s.length
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if length < 0 {
		return nil, fmt.Errorf("Length must be greater than 0")
	}

	square := &Square{
		center: Point{
			x: x,
			y: y,
		},
		length: length,
	}

	return square, nil
}

func main() {
	s, err := NewSquare(1, 2, 3)
	if err != nil {
		fmt.Printf("Error: %s \n", err)
		os.Exit(0)
	}
	fmt.Println(s.center)
	fmt.Println(s.Area())
	s.Move(10, -10)
	fmt.Println(s.center)
}
