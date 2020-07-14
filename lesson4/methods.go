package main

import "fmt"

type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}
	return value
}

func main() {
	t1 := Trade{"MSFT", 10, 99.8, true}
	fmt.Println(t1.Value())
}
