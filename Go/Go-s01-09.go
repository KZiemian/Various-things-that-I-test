package main

import "fmt"

var c, python, java bool

func main() {
	var i int

	fmt.Printf("i: %v, %#v, %T\n", i, i, i)
	fmt.Printf("c: %v, %#v, %T\n", c, c, c)
	fmt.Printf("java: %v, %#v, %T\n", java, java, java)
}
