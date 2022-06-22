package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	fmt.Printf("v1: %v, %#v, %T\n", v1, v1, v1)
	fmt.Printf("p: %v, %#v, %T\n", p, p, p)
	fmt.Printf("v2: %v, %#v, %T\n", v2, v2, v2)
	fmt.Printf("v3: %v, %#v, %T\n", v3, v3, v3)
}
