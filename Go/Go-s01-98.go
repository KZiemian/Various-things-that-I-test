package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	var a Student

	fmt.Printf("a: %v, %#v, %T\n\n", a, a, a)

	a.Name = "Alice"
	fmt.Printf("a: %v, %#v, %T\n", a, a, a)
}
