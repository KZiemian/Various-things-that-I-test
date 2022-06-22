package main

import "fmt"

func Min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func main() {
	fmt.Printf("Min(%v, %v) -> %v\n", 0, 0, Min(0, 0))
	fmt.Printf("Min(%v, %v) -> %v\n", 0, 1, Min(0, 1))
	fmt.Printf("Min(%v, %v) -> %v\n", 1, 0, Min(1, 0))
	fmt.Printf("Min(%v, %v) -> %v\n", 1, 2, Min(1, 2))
}
