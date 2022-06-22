package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Printf("f.Abs(): %v\n\n", f.Abs())

	fmt.Printf("math.Sqrt2: %v, %#v, %T\n", math.Sqrt2,
		math.Sqrt2, math.Sqrt2)
	fmt.Printf("f: %v, %#v, %T\n", f, f, f)
}
