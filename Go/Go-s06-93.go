package main

import "fmt"

const MaxUint = ^uint(0)

func main() {
	fmt.Printf("MaXUint: %v, %T\n", MaxUint, MaxUint)

	var f64 float64 = MaxUint
	fmt.Printf("f64: %v, %T\n", f64, f64)
}
