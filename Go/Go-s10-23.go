package main

import (
	"fmt"
	"math"
)

func main() {
	x := 0.0
	frac, exp := math.Frexp(x)

	fmt.Printf("math.Frexp(%v) = %v, %v\n", x, frac, exp)

	x = 1.0
	frac, exp = math.Frexp(x)
	fmt.Printf("math.Frexp(%v) = %v, %v\n", x, frac, exp)

	x = 2.0
	frac, exp = math.Frexp(x)
	fmt.Printf("math.Frexp(%v) = %v, %v\n", x, frac, exp)

	x = 3.0
	frac, exp = math.Frexp(x)
	fmt.Printf("math.Frexp(%v) = %v, %v\n", x, frac, exp)

	x = 4.0
	frac, exp = math.Frexp(x)
	fmt.Printf("math.Frexp(%v) = %v, %v\n", x, frac, exp)

	x = 5.0
	frac, exp = math.Frexp(x)
	fmt.Printf("math.Frexp(%v) = %v, %v\n", x, frac, exp)

	x = 16.0
	frac, exp = math.Frexp(x)
	fmt.Printf("math.Frexp(%v) = %v, %v\n", x, frac, exp)
}
