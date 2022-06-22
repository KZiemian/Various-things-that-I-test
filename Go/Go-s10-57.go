package main

import (
	"fmt"
	"math"
)

func main() {
	for x := 0.05; x <= 5; x += 0.05 {
		fmt.Printf("math.Yn(2, %.2f) = %v\n", x, math.Yn(2, x))
	}
}
