package main

import "fmt"

type Student struct {
	Name string
}

func main() {
	ps := new(Student)
	fmt.Printf("ps: %v, %#v, %T\n\n", ps, ps, ps)

	ps.Name = "Alice"
	fmt.Printf("ps: %v, %#v, %T\n", ps, ps, ps)
}
