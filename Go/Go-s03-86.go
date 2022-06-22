package main

import "fmt"

func main() {
	a := [2]string{"Penn", "Teller"}
	b := [...]string{"Penn", "Teller"}
	c := []string{"Penn", "Teller"}

	fmt.Printf("a: %v, %T\n", a, a)
	fmt.Printf("b: %v, %T\n", b, b)
	fmt.Printf("c: %v, %T\n", c, c)
}
