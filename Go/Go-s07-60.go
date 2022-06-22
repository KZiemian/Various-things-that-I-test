package main

import "fmt"

type Foo struct {
	i int
	s string
}

func main() {
	var s = []Foo{
		{6 * 9, "Question"},
		{42, "Answer"},
	}

	var m = map[int]Foo{
		7: {6 * 9, "Question"},
		3: {42, "Answer"},
	}

	fmt.Printf("s: %v, %T\n", s, s)
	fmt.Printf("m: %v, %T\n", m, m)
}
