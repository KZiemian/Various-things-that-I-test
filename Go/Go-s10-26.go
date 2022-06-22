package main

import (
	"fmt"
	"math"
)

func main() {
	for x := 0.0; x <= 17; x += 1.0 {
		fmt.Printf("math.Ilogb(%.4f) = %v\n", x, math.Ilogb(x))
	}
}
