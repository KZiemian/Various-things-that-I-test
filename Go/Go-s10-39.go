package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("math.Min(0, 0) = %v\n", math.Min(0, 0))
	fmt.Printf("math.Min(1, 0) = %v\n", math.Min(1, 0))
	fmt.Printf("math.Min(1, -1) = %v\n", math.Min(1, -1))
}
