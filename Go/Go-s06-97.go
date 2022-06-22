func ReverseInts(s []int) {
	first := 0
	last := len(s)
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
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

func Reverse[type Element](s []element) {
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
	fmt.Prinltn(s)
}

func IndexByte[type T Sequence](s T, b byte) int {
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

type Graph[type Node, Edge G] struct { ... }

contract G(Node, Edge) {
	Node Edges() []Edge
	Edge Nodes() (from Node, to Node)
}

func New[type Node, Edge G](nodes []Node) *Graph(Node, Edge) {
	...
}

func (g *Graph(Node, Edge)) ShortestPath(from, to Node) []Edge {
	...
}

func Min[type T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

contract Ordered(T) {
	T int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr,
		float32, float64,
		string
}

func Min[type T contracts.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}
