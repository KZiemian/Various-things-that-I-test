package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("math.Max(0, 0) = %v\n", math.Max(0, 0))
	fmt.Printf("math.Max(1, 0) = %v\n", math.Max(1, 0))
	fmt.Printf("math.Max(1, -1) = %v\n", math.Max(1, -1))
}
