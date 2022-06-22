package main

import "fmt"

func main() {
	var s []byte
	s = make([]byte, 5, 5)

	fmt.Printf("s: %v, %T\n", s, s)

	sOne := make([]byte, 5)
	fmt.Printf("s: %v, %T\n", sOne, sOne)

	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	c := b[1:4]

	fmt.Printf("b: %v, %T\n", b, b)
	fmt.Printf("c: %v, %T\n", c, c)
}
