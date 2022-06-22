package main

import "fmt"

type NodeConstraint[2] interface {
	Edges() []int
}

type NodeInt int

func main() {
	sliceNodes := []NodeInt{NodeInt(0), NodeInt(1), NodeInt(2)}

	fmt.Println("sliceNodes:", sliceNodes)
}
