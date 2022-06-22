package main

import "fmt"

func main() {
	a := make([]int, 1)
	fmt.Println("a:", a)
	fmt.Printf("a: %v, %#v, %T\n", a, a, a)
	fmt.Printf("len(a): %v, cap(a): %v\n", len(a), cap(a))

	a = append(a, 1, 2, 3)
	fmt.Println("a:", a)
	fmt.Printf("a: %v, %#v, %T\n", a, a, a)
	fmt.Printf("len(a): %v, cap(a): %v\n", len(a), cap(a))
}
