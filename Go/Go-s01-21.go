package main

import "fmt"

func main() {
	var i int = 42
	j := i
	f := 3.142
	g := 0.867 + 0.5i

	fmt.Printf("i: %v, %#v, %T\n", i, i, i)
	fmt.Printf("j: %v, %#v, %T\n", j, j, j)
	fmt.Printf("f: %v, %#v, %T\n", f, f, f)
	fmt.Printf("g: %v, %#v, %T\n", g, g, g)
}
