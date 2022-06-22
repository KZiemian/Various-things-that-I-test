package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.3f\n", math.FMA(1, 2, 3))
	fmt.Printf("%.3f\n", math.FMA(2, 3, 1))
	fmt.Printf("%.3f\n", math.FMA(3, 4, 1))
}
