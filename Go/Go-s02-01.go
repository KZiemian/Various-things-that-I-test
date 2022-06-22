package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	b := Student{
		Name: "Bob",
	}

	fmt.Printf("b: %v, %#v, %T\n\n", b, b, b)

	pb := &Student{
		Name: "Bob",
		Age:  8,
	}

	fmt.Printf("pb: %v, %#v, %T\n", pb, pb, pb)
	fmt.Printf("*pb: %v, %#v, %T\n\n", *pb, *pb, *pb)

	c := Student{"Cecilia", 5}
	fmt.Printf("c: %v, %#v, %T\n\n", c, c, c)

	d := Student{}
	fmt.Printf("d: %v, %#v, %T\n", d, d, d)
}
