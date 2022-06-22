package main

import "fmt"

type Foo struct {
	i int
	s string
}

type Bar int

type Qux []string

func main() {
	var i int
	var s []string

	var t struct {
		i int
		s string
	}

	fmt.Printf("i: %v, %T\n", i, i)
	fmt.Printf("s: %v, %T\n", s, s)
	fmt.Printf("t: %v, %T\n", t, t)
}
