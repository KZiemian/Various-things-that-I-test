package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Printf("add(%v, %v) = %v\n", 42, 13, add(42, 13))
}
