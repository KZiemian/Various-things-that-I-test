package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v

	fmt.Printf("v: %v, %#v, %T\n", v, v, v)
	fmt.Printf("p: %v, %#v, %T\n", p, p, p)
	fmt.Printf("p.X: %v, %#v, %T\n", p.X, p.X, p.X)

	p.X = 1e9
	fmt.Printf("v: %v, %#v, %T\n", v, v, v)
	fmt.Printf("p: %v, %#v, %T\n", p, p, p)
	fmt.Printf("p.X: %v, %#v, %T\n", p.X, p.X, p.X)
}
