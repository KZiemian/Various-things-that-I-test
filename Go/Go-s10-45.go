package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("math.Remainder(0, 1) = %v\n", math.Remainder(0, 1))
	fmt.Printf("math.Remainder(1, 2) = %v\n", math.Remainder(1, 2))
	fmt.Printf("math.Remainder(2, 2) = %v\n", math.Remainder(2, 2))
	fmt.Printf("math.Remainder(3, 2) = %v\n", math.Remainder(3, 2))
}
