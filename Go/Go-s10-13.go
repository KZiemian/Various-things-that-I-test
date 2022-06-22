package main

import (
	"fmt"
	"math"
)

func main() {
	for x := 0.0; x <= 2.05; x += 0.1 {
		fmt.Printf("math.Erfcinv(%.1f) = %v\n", x, math.Erfcinv(x))
	}
}
