package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	var v Vertex = Vertex{1, 2}

	fmt.Printf("v: %v, %#v, %T\n\n", v, v, v)
	fmt.Printf("v.X: %v, %#v, %T\n", v.X, v.X, v.X)
	fmt.Printf("v.Y: %v, %#v, %T\n\n", v.Y, v.Y, v.Y)

	v.X = 4
	fmt.Printf("v: %v, %#v, %T\n", v, v, v)
	fmt.Printf("v.X: %v, %#v, %T\n", v.X, v.X, v.X)
	fmt.Printf("v.Y: %v, %#v, %T\n", v.Y, v.Y, v.Y)
}
