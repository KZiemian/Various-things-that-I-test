package main

import "fmt"

type Student struct {
	Name string
}

func main() {
	var ps1 *Student = new(Student)
	fmt.Printf("ps1: %v, %#v, %T\n\n", ps1, ps1, ps1)

	ps2 := new(Student)
	fmt.Printf("ps2: %v, %#v, %T\n", ps2, ps2, ps2)
}
