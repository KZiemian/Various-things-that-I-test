package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"

	fmt.Printf("i: %v, %#v, %T\n", i, i, i)
	fmt.Printf("j: %v, %#v, %T\n", j, j, j)
	fmt.Printf("c: %v, %#v, %T\n", c, c, c)
	fmt.Printf("python: %v, %#v, %T\n", python, python, python)
	fmt.Printf("java: %v, %#v, %T\n", java, java, java)
}
