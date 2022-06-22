type Tree [type E] struct {
	root    *node(E)
	compare func(E, E) int
}

type node [type E] struct {
	val         E
	left, right *node(E)
}

func New[type E](cmp func(E, E) int) *Tree(E) {
	return &Tree[E]{compare: cmp}
}

func (t *Tree(E)) find(v E) **node(E) {
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

func (t *Tree(E)) Contains(v E) bool {
	return *t.find(e) != nil
}

func (t *Tree(E)) Insert(v E) bool {
	pn := t.find(v)
	if *pn != nil {
		return false
	}
	*pn = &node(E){val: v}
	return true
}

var intTree = tree.New(func(a, b int) int { return a - b })

func InsertAndCheck(v int) {
	intTree.Insert(v)
	if !intTree.Contains(v) {
		log.Fatalf("%d not found after insertion", v)
	}
}
