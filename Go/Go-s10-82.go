func (t *Tree[E]) Find(v E) bool {
	return *t.find(e) != nil
}

func (t *Tree[E]) Insert(v E) bool {
	pn := t.find(v)
	if *pn != nil {
		return false
	}
	*pn = &node[E]{val: V}

	return true
}

var intTree = tree.New(
	func(a, b int) int { return a -  b } )

func InsertAndCheck(v int) {
	intTree.Insert(v)
	if !intTree.Fing(v) {
		log.Fatalf("%d not found after insertion", v)
	}
}

func ReverseInts(s []int) {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}

func ReverseStrings(s []string) {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}

func Reverse[type Element](s []Element) {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}

func ReverseAndPrint(s []int) {
	Reverse[int](s)
	fmt.Println(s)
}

func ReverseAndPrint(s []int) {
	Reverse(s)
	fmt.Println(s)
}

func IndexByte[type T Sequence](s T, b byte) {
	for i := 0; i < len(s); i++ {
		if s[i] == b {
			return i
		}
	}
	return -1
}

contract Sequence(T) {
	T string, []byte
}

func ToStrings[type E Stringer](s []E) []string {
	r := make([]string, len(s))
	for i, v := range s {
		r[i] = v.String()
	}

	return r
}

contract Stringer(T) {
	T String() string
}

type Graph[type Node, Edge G] struct {
	// ...
}

contract G(Node, Edge) {
	Node Edges() []Edge
	Edge Nodes() (from Node, to Node)
}

func New[type Node, Edge G](nodes []Node) *Graph(Node, Edge) {
	// ...
}

func (g *Graph(Node, Edge)) ShortestPath(from, to Node) []Edge {
	// ...
}
