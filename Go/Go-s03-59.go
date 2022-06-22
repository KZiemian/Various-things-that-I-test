package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	for i := 1; i < 10; i++ {
		z -= (z*z - x) / (2*z)

		fmt.Printf("i: %v, z: %v\n", i, z)
	}

	return z
}

func main() {
	z := 40.0

	fmt.Printf("     Sqrt(%v): %g\n", z, Sqrt(z))
	fmt.Printf("math.Sqrt(%v): %g\n", z, math.Sqrt(z))
}
