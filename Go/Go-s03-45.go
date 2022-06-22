package main

import "fmt"

func main() {
	var i int
	j := i

	fmt.Printf("i: %v, %T\n", i, i)
	fmt.Printf("j: %v, %T\n", j, j)
}
