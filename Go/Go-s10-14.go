package main

import (
	"fmt"
	"math"
)

func main() {
	for x := -1.0; x <= 1.0; x += 0.1 {
		fmt.Printf("math.Erfinv(%.2f) = %v\n", x, math.Erfinv(x))
	}

	fmt.Printf("math.Erfinv(0.99999999999999) = %v\n",
		math.Erfinv(0.99999999999999))
	fmt.Printf("math.Erfinv(1) = %v\n", math.Erfinv(1))
}
