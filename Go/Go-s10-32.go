package main

import (
	"fmt"
	"math"
)

func main() {
	lgamma, sign := math.Lgamma(1)

	for x := 1.0; x <= 10; x += 0.5 {
		lgamma, sign = math.Lgamma(x)

		fmt.Printf("math.Gamma(%.3f) = %v\n", x, math.Gamma(x))
		fmt.Printf("math.Lgamma(%.3f) = %v, %v\n", x,
			lgamma, sign)
	}
}
