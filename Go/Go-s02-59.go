package main

import "fmt"

func multiplyByTwo(x int) (y int) {
	defer func() { y = 2*y }()
	return x
}

func main() {
	for i := 0; i < 4; i++ {
		fmt.Printf("multiplyByTwo(%v) = %v\n", i, multiplyByTwo(i))
	}
}
