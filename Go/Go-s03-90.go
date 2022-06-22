package main

import "fmt"

func main() {
	d := []byte{'r', 'o', 'a', 'd'}
	e := d[2:]

	fmt.Printf("d: %v, %T\n", d, d)
	fmt.Printf("e: %v, %T\n", e, e)

	e[1] = 'm'
	fmt.Printf("e: %v, %T\n", e, e)
	fmt.Printf("d: %v, %T\n", d, d)
}
