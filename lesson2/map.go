package main

import "fmt"

func main() {
	stonks := map[string]float64{
		"AMZN": 111,
		"GOOG": 222,
		"MSFT": 333,
	}

	fmt.Println(len(stonks))

	fmt.Println(stonks["MSFT"])

	fmt.Println(stonks["TSLA"])

	value, ok := stonks["TSLA"]
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}

	stonks["TSLA"] = 444
	fmt.Println(stonks)

	delete(stonks, "AMZN")
	fmt.Println(stonks)

	for key := range stonks {
		fmt.Println(key)
	}

	for key, value := range stonks {
		fmt.Println(key)
		fmt.Println(value)
	}
}
