package main

import (
	"fmt"
	// "math"
)

type Constraint interface {
	float64
}

func someFunction[uint8 Constraint](x uint8) uint8 {
	x += uint8(1)
	return x
}

func main() {
	var x float64 = 1

	fmt.Printf("someFunction(%v): %v\n", x, someFunction(x))
	fmt.Printf("x type: %T\n", x)
	fmt.Printf("someFunction(x) type: %T\n", someFunction(x))
}
