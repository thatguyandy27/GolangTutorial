package main

import (
	"fmt"
	"math"
)

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("sqrt of negative value - %f", n)
	}

	return math.Sqrt(n), nil
}

func main() {
	s1, err := sqrt(64)

	s2, err := sqrt(-64)
}
