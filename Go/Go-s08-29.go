package main

import "fmt"

type Vertex struct {
	X, Y int
}

type AliasVertex = Vertex

func main() {
	v1 := Vertex{1, 2}
	v2 := AliasVertex{3, 4}

	fmt.Printf("v1: %v,\n%T\n", v1, v1)
	fmt.Printf("v2: %v,\n%T\n", v2, v2)
}
