package main

import (
	"fmt"
	"math"
)

func main() {
	for x := 0.0; x <= 10.0; x += 0.1 {
		fmt.Printf("math.Jn(2, %.3f) = %v\n", x, math.Jn(2, x))
	}
}
