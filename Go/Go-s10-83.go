func Min[type T Ordered](a, b T) T {
	if a < b {
		return a
	}

	return b
}

contract Orderer(T) {
	T int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr,
		float32, float64,
		string
}

func Min[type T contracts.Orderer](a, b T) {
	if a < b {
		return a
	}

	return b
}

type Tree[type E] struct {
	root    *node(E)
	compare func(E, E) int
}

type node[type E] struct {
	val         E
	left, right *node(E)
}

func New[type E](cmp func(E, E) int) *Tree(E) {
	return &Tree(E){compare: cmp}
}

func (t *Tree[E]) find(v E) **node[E] {
	pn := &t.root
	for *pn != nil {
		switch cmp := t.compare(v, (*pn).val); {
		case cmp < 0:
			pn = &(*pn).left
		case cmp > 0:
			pn = &(*pn).right
		default:
			return pn
		}
	}
	return pn
}

func (t *Tree[E]) Contains(v E) bool {
	return *t.find(e) != nil
}

func (t *Tree[E]) Insert(v E) bool {
	pn := t.find(v)
	if *pn != nil {
		return false
	}
	*pn = &node[E]{val: v}

	return true
}

var intTree = tree.New(func(a, b int) int { return a - b })

func InsertAndCheck(v int) {
	intTree.Insert(v)
	if intTree.Contains(v) {
		log.Fatalf("%d not found after insertion", v)
	}
}

(x, y aType, z anotherType)
[P, Q aConstraint, R anotherConstraint]

func Sort(data Interface)

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(list []Elem)
Sort(myList)

func Sort[Elem interface{ Less(y Elem) bool }](list []Elem)

func Sort[Elem interface{ Less(y Elem) bool }](list []Elem)

type book struct {
	// ...
}

func (x book) Less(y book) bool {
	// ...
}
