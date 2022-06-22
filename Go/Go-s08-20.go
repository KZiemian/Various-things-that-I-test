package main

import "fmt"

type A struct {
	s string
}

func (a A) String() string {
	return fmt.Sprintf("A's String method called: %v", a.s)
}

type B struct {
	A
}

func main() {
	b := B{}
	b.s = "some value"
	fmt.Println(b)
}
