package main

import "fmt"

func main() {
	var t struct {
		i int
		s string
	}

	t.i = 3
	t.s = "Martin"

	fmt.Printf("t: %v, %T\n", t, t)
}
