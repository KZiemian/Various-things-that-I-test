type MyString string

type AnyString interface {
	~string
}

type ApproximateMyString interface {
	~MyString // INVALID: underlying type of MyString is not MyString
}

type ApproximateParameter[T any] interface {
	~T // INVALID: T is a type parameter
}

type PredeclaredSignedInteger interface {
	int | int8 | int16 | int32 | int64
}

type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

package constraints

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func Smallest[T constraints.Ordered](s []T) T {
	r := s[0]
	for _, v := range s[1:] {
		if v < r {
			r = v
		}
	}
	return r
}

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

type ComparableHasher interface {
	comparable
	Hash() uintptr
}

type ImpossibleConstraint interface {
	comparable
	[]int
}

package graph

type NodeConstraint[Edge any] interface {
	Edges() []Edge
}

type EdgeConstraing[Node any] interface {
	Nodes() (from, to Node)
}

type Graph[Node NodeConstraint[Edge], Edge EdgeConstraint[Node]] struct {
	// ...
}

func New[Node NodeConstraint[Edge], Edge EdgeConstraint[Node]](nodes []Node) *Graph[Node, Edge] {
	// ...
}

func (g *Graph[Node, Edge]) ShortestPath(from, to Node) []Edge {
	// ...
}

type Vertex struct {
	// ...
}

func (v *Vertex) Edges() []*FromTo {
	// ...
}

type FromTo struct {
	// ...
}

func (ft *FromTo) Nodes() (*Vertex, *Vertex) {
	// ...
}

var g = graph.New[*Vertex, *FromTo]([]*Vertex{ ... })

type NodeInterface interface { Edges() []EdgeInterface }

type EdgeInterface interface { Nodes() (NodeInterface, NodeInterface) }
