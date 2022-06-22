package main

import (
	"fmt"
	"math"
)

func main() {
	exp := 2

	for x := 0.0; x <= 10; x += 1.0 {
		fmt.Printf("math.Ldexp(%.3f, %d) = %v\n", x, exp,
			math.Ldexp(x, exp))
	}
}
