package main

import "fmt"

func main() {
	i := 42
	f := 3.142
	g := 0.867 + 0.5i

	fmt.Printf("i: %v, %T\n", i, i)
	fmt.Printf("f: %v, %T\n", f, f)
	fmt.Printf("g: %v, %T\n", g, g)
}
