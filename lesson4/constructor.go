package main

import (
	"fmt"
	"os"
)

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

func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("Symbol cannot be null")
	}
	if volume <= 0 {
		return nil, fmt.Errorf("Volume must be greater than 0")
	}
	if price <= 0 {
		return nil, fmt.Errorf("Price must be greater than 0")
	}
	trade := &Trade{
		Symbol: symbol,
		Volume: volume,
		Price:  price,
		Buy:    buy,
	}
	return trade, nil
}

func main() {
	t, err := NewTrade("MSFT", -100, 12, true)
	if err != nil {
		fmt.Printf("Error, cannot create trade: %s\n", err)
		os.Exit(0)
	} else {
		fmt.Println(t.Value())
	}
}
