package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 1; i < 10; i++ {
		fmt.Printf("math.Logb(%d) = %.2f\n", i,
			math.Logb(float64(i)))
	}
}
