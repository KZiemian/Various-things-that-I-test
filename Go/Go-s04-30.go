package main

import "fmt"

func main() {
	var nasdaq chan struct {
		symbol string
		price  float64
	}

	m := map[string]float64{}

	for x := range nasdaq {
		last := m[x.symbol]

		if x.price > last {
			fmt.Printf("buy %s!\n", x.symbol)
		}

		if x.price < last {
			fmt.Printf("sell %s\n", x.symbol)
		}

		m[x.symbol] = x.price
	}
}
