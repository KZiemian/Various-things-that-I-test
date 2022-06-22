package main

import (
	"fmt"
	"math"
)

func main() {
	for x := 0.05; x <= 5; x += 0.05 {
		fmt.Printf("math.Y1(%.2f) = %v\n", x, math.Y1(x))
	}
}
