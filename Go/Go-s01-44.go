package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println("hypot(5, 12):", hypot(5, 12))
	fmt.Printf("hypot: %v, %#v, %T\n", hypot, hypot, hypot)

	fmt.Printf("math.Sqrt: %v, %#v, %T\n", math.Sqrt,
		math.Sqrt, math.Sqrt)
}
