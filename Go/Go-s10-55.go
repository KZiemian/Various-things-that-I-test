package main

import (
	"fmt"
	"math"
)

func main() {
	for x := 0.1; x <= 10; x += 0.1 {
		fmt.Printf("math.Y0(%.2f) = %v\n", x, math.Y0(x))
	}
}
