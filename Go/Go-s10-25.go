package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("math.Hypot(0, 0) = %.4f\n", math.Hypot(0, 0))
	fmt.Printf("math.Hypot(1, 0) = %.4f\n", math.Hypot(1, 0))
	fmt.Printf("math.Hypot(0, 1) = %.4f\n", math.Hypot(0, 1))
	fmt.Printf("math.Hypot(2, 0) = %.4f\n", math.Hypot(2, 0))
	fmt.Printf("math.Hypot(0, 2) = %.4f\n", math.Hypot(0, 2))
	fmt.Printf("math.Hypot(2, 1) = %.4f\n", math.Hypot(2, 1))
	fmt.Printf("math.Hypot(1, 2) = %.4f\n", math.Hypot(1, 2))
	fmt.Printf("math.Hypot(1.4142, 1.4142) = %.4f\n",
		math.Hypot(1.4142, 1.4142))

}
