package main

import "fmt"

func main() {
	var x, y, max, min int = 2, 1, 0, 0

	fmt.Printf("x: %v; y: %v; max: %v; min: %v\n", x, y, max, min)

	if x > max {
		max = x
	}
	fmt.Printf("x: %v; y: %v; max: %v; min: %v\n", x, y, max, min)

	if x <= y {
		min = x
	} else {
		min = y
	}
	fmt.Printf("x: %v; y: %v; max: %v; min: %v\n", x, y, max, min)
}
