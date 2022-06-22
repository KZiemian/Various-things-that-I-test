package main

import (
	"fmt"
	"math"
)

func main() {
	for x := 0.0; x < 10; x += 0.1 {
		fmt.Printf("math.Erf(%.1f): %v\n", x, math.Erf(x))
	}
}
