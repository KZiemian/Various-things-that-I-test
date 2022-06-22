package main

import (
	"fmt"
	"math"
)

func main() {
	f1 := float64(64)
	bits1 := math.Float64bits(f1)
	f2 := math.Float64frombits(bits1)

	fmt.Printf("f1: %v\n", f1)
	fmt.Printf("bits1: %v, %b\n", bits1, bits1)
	fmt.Printf("f1 == f2: %v\n", f1 == f2)
}
