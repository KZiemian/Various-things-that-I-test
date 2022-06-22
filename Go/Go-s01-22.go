package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	varVertex := Vertex{1, 2}

	fmt.Printf("varVertex: %v, %#v, %T\n", varVertex, varVertex,
		varVertex)
	fmt.Printf("Vertex{%v, %v}: %v, %#v, %T\n", 1, 2, Vertex{1, 2},
		Vertex{1, 2}, Vertex{1, 2})
	fmt.Printf("Vertex{%v, %v}: %v, %#v, %T\n", 1, 2,
		Vertex{1, 2}, Vertex{1, 2}, Vertex{1, 2})
}
