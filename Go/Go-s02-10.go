package main

import "fmt"

func f() int {
	return 1
}

func main() {
	var y int = 2

	if x := f(); x <= y {
		fmt.Printf("x: %v\n", x)
	} else {
		fmt.Printf("y: %v\n", y)
	}

	// fmt.Printf("x: %v\n", x)
}
