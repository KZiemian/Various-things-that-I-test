package main

import "fmt"

func funcTest[T int | float64](x T) T {
	var v T

	v = 2 * x

	return v
}

func main() {
	fmt.Printf("funTest(2) = %v\n", funcTest(2))
}
