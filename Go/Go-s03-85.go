package main

import "fmt"

func main() {
	var a [4]int

	fmt.Printf("a: %v, %T\n", a, a)
	fmt.Printf("len(a): %v, %T\n", len(a), len(a))
	fmt.Printf("cap(a): %v, %T\n", cap(a), cap(a))

	a[0] = 1
	i := a[0]

	fmt.Printf("a: %v, %T\n", a, a)
	fmt.Printf("a: %v, %T\n", i, i)
}
