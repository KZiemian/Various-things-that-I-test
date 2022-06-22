package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x
	zPrev := z

	for i := 1; i < 10; i++ {
		zPrev = z
		z -= (z*z - x) / (2*z)

		fmt.Printf("i: %v, z: %v\n", i, z)

		if math.Abs(z - zPrev) < 0.0000001 {
			break
		}
	}

	return z
}

func main() {
	z := 10000.0

	fmt.Printf("     Sqrt(%v): %v\n", z, Sqrt(z))
	fmt.Printf("math.Sqrt(%v): %v\n", z, math.Sqrt(z))
}
