package main

import (
	"fmt"
	"math"
)

func main() {
	for x := 0.0; x <= 10; x += 0.1 {
		fmt.Printf("math.J1(%.3f) = %v\n", x, math.J1(x))
	}
}
