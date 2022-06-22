package main

import (
	"fmt"
	"math"
)

func main() {
	for x := 0.0; x <= 10.0; x += 0.1 {
		fmt.Printf("J0(%.3f) = %v\n", x, math.J0(x))
	}
}
