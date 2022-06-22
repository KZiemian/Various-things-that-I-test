package main

import (
	"fmt"
	"math"
)

func main() {
	var v float32 = 13
	x := math.Float32bits(v)
	fmt.Printf("%v; %b\n", x, x)
}
