package main

import "fmt"

func someFunction(range int) int {
	return 2*range
}

func main() {
	fmt.Printf("someFunction(2): %v\n", someFunction(2))
}
