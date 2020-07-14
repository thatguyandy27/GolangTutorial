package main

import (
	"fmt"
	"math"
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
	length float64
}

func (s *Square) Move(dx int, dy int) {
	s.center.Move(dx, dy)
}

func (s *Square) Area() float64 {
	return s.length * s.length
}

func NewSquare(x int, y int, length float64) (*Square, error) {
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

type Circle struct {
	center Point
	radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func sumAreas(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}

	return total
}

type Shape interface {
	Area() float64
}

func main() {
	s := &Square{center: Point{0, 0}, length: 10.0}
	c := &Circle{center: Point{0, 0}, radius: 5.0}

	shapes := []Shape{s, c}
	total := sumAreas(shapes)
	fmt.Println(total)
}
