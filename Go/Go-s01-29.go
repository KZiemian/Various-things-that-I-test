package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ring",
	}
	fmt.Printf("names: %v, %#v, %T\n\n", names, names, names)

	a := names[0:2]
	b := names[1:3]
	fmt.Printf("a: %v, %#v, %T\n", a, a, a)
	fmt.Printf("b: %v, %#v, %T\n\n", b, b, b)

	b[0] = "XXX"
	fmt.Printf("a: %v, %#v, %T\n", a, a, a)
	fmt.Printf("b: %v, %#v, %T\n", b, b, b)
	fmt.Printf("names: %v, %#v, %T\n", names, names, names)
}
