package main

import "fmt"

func c(x int) (y int) {
	defer func() { y = 2*y }()
	return x
}

func main() {
	fmt.Printf("c(%v) = %v\n", 1, c(1))
	fmt.Printf("c(%v) = %v\n", 2, c(2))
}
