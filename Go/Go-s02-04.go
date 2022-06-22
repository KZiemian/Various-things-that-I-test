package main

import "fmt"

type Student struct {
	Name string
}

func main() {
	s := Student{"Alice"}
	ps := &s

	fmt.Printf("s: %v, %#v, %T\n", s, s, s)
	fmt.Printf("ps: %v, %#v, %T\n", ps, ps, ps)

	ps = &Student{"Alice"}
	fmt.Printf("ps: %v, %#v, %T\n", ps, ps, ps)
}
