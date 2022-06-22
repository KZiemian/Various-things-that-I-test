package main

import (
	"fmt"
	"math"
)

func main() {
	for x := 0.0; x < 10; x += 0.1 {
		fmt.Printf("math.Erfc(%.1f) = %v\n", x, math.Erfc(x))
	}
}
