package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	var pa *Student

	fmt.Printf("pa: %v, %#v, %T\n\n", pa, pa, pa)
	// fmt.Printf("*pa: %v, %#v, %T\n\n", *pa, *pa, *pa)

	pa = new(Student)
	fmt.Printf("pa: %v, %#v, %T\n", pa, pa, pa)
	fmt.Printf("*pa: %v, %#v, %T\n", *pa, *pa, *pa)
	pa.Name = "Alice"
	fmt.Printf("pa: %v, %#v, %T\n", pa, pa, pa)
	fmt.Printf("*pa: %v, %#v, %T\n", *pa, *pa, *pa)
}
