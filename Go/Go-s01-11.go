package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Printf("i: %v, %#v, %T\n", i, i, i)
	fmt.Printf("j: %v, %#v, %T\n", j, j, j)
	fmt.Printf("k: %v, %#v, %T\n", k, k, k)
	fmt.Printf("c: %v, %#v, %T\n", c, c, c)
	fmt.Printf("python: %v, %#v, %T\n", python, python, python)
	fmt.Printf("java: %v, %#v, %T\n", java, java, java)
}
