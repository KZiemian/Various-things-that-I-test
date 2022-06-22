package main

import "fmt"

type NodeConstraint[Edge any] interface {
	Edges() []Edge
}

type NodeChain int

func (node NodeChain) Edges() []EdgeChain {
	return []EdgeChain{EdgeChain{int(node) - 1, int(node)},
		EdgeChain{int(node), int(node) + 1},
	}
}

type EdgeConstraint[Node any] interface {
	Nodes() (from, to Node)
}

type EdgeChain struct {
	From int
	To   int
}

func (edge EdgeChain) Nodes() (from, to NodeChain) {
	return NodeChain(edge.From), NodeChain(edge.To)
}

type Graph[Node NodeConstraint[Edge], Edge EdgeConstraint[Node]] struct {
	NodesSlice []Node
	EdgesSlice []Edge
}

func New[Node NodeConstraint[Edge], Edge EdgeConstraint[Node]](nodes []Node) *Graph[Node, Edge] {
	var edges []Edge

	for _, node := range nodes {
		for _, edge := range node.Edges() {
			edges = append(edges, edge)
		}
	}
	return &Graph[Node, Edge]{nodes, edges}
}

func main() {
	sliceNodes := []NodeChain{NodeChain(1), NodeChain(2), NodeChain(3)}

	ptrGraph := New[NodeChain, EdgeChain](sliceNodes)

	fmt.Printf("sliceNodes: %v, %T\n", sliceNodes, sliceNodes)
	fmt.Printf("ptrGraph: %v, %T\n", ptrGraph, ptrGraph)

	// // fmt.Println("sliceNodes:", sliceNodes)


	// node := NodeChain(1)

	// fmt.Printf("node value: %v\n", node)
	// fmt.Printf("node type:  %T\n\n", node)

	// fmt.Printf("node.Edges() result value: %v\n", node.Edges())
	// fmt.Printf("node.Edges() result type:  %T\n\n", node.Edges())


	// edge := EdgeChain{7, 8}

	// fmt.Printf("edge value: %v\n", edge)
	// fmt.Printf("edge type:  %T\n", edge)

	// from, to := edge.Nodes()
	// fmt.Printf("edge.Nodes() value: %v, %v\n", from, to)
	// fmt.Printf("edge.Nodes() type:  %T, %T\n", from, to)
}
