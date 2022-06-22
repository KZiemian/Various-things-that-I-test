package main

import (
	"fmt"
	"math"
)

func main() {
	for x := 1.0; x < 3; x += 0.1 {
		fmt.Printf("math.Gamma(%.3f) = %v\n", x, math.Gamma(x))
	}
}
