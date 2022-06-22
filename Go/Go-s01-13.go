package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("i: %v, %#v, %T\n", i, i, i)
	fmt.Printf("f: %v, %#v, %T\n", f, f, f)
	fmt.Printf("b: %v, %#v, %T\n", b, b, b)
	fmt.Printf("s: %v, %#v, %T\n", s, s, s)
}
