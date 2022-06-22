package main

import (
	"fmt"
	"math"
)

func main() {
	u := math.Float32bits(64)
	f := math.Float32frombits(u)

	fmt.Printf("u: %v, %b\n", u, u)
	fmt.Printf("f: %v, %b\n", f, f)
	fmt.Printf("float32(64) == f: %v\n", float32(64) == f)
}
