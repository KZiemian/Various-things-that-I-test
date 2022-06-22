func (g *Graph[Node, Edge]) ShortestPath(from, to Node) []Edge {
	// ...
}

func MapKeys(m map[string]int) []string {
	var s []string

	for k := range m {
		s = append(s, k)
	}

	return s
}

package graph

type NodeConstraint[Edge any] interaface {
	Edges() []Edge
}

type EdgeConstraint[Node any] interface {
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

func Map[F, T any](s []F, f func(F) T) []T {
	// ...
}

var s []int

f := func(i int) int 64 { return int64(i) }

var r []int64

r = Map[int, int64](s, f)

r = Map[int](s, f)

r = Map(s, f)

[]map[int]bool

[]map[int]bool

T1 // T1 == map[int]bool

[]T1 // T1 == map[int]bool

[]map[T1]T2 // T1 == int, T2 == bool

Print[int]([]int{1, 2, 3})

s1 := []int{1, 2, 3}

Print(s1)

func Map[F, T any](s []F, f func(F) T) []T {
	r := make([]T, len(s))

	for i, v := range s {
		r[i] = f(v)
	}

	return r
}

strs := Map([]int{1, 2, 3}, strconv.Itoa)

func NewPair[F any](f1, f2 F) *Pair[F] { ... }

NewPair(1, 2)

NewPair[int](1, 2)
