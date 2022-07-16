type C interface {
	int
	M()
}

func F(x, y, z bool) bool {
	return ((!x && y) || z) && (x || !y)
}

((!true && true) || false) && (true || !true)
((false && true) || false) && (true || !true)
((false && true) || false) && (true || false)
((false && true) || false) && true
(false && true) || false
false && true
false

((!x && y) && (x || !y)) || (z && (x || !y))
(((!x && y) && x) || ((!x && y) && !y)) || (z && x) || (z && !y)
false || false || (z && x) || (z && !y)
(z && x) || (z && !y)
(x && z) || (!y && z)

func F_DNF(x, y, z bool) bool {
	return (x && z) || (!y && z)
}

((!x && y) || z) && (x || y)
((!x || z) && (y || z)) && (x || y)
(!x || z) && (y || z) && (x || y)

var Univers Set[T]

func MakeSet(f func(T) bool) Set[T] {
	s := make(Set[T])

	for v := range Universe {
		if f(v) {
			s.Add(v)
		}
	}

	return s
}

func MakeFunc(s Set[T]) func(T) bool {
	return func(v T) bool {
		return s.Contains(v)
	}
}

func Union(s, t Set[T]) Set[T] {
	return MakeSet(func(v T) bool {
		return s.Contains(v) || t.Contains(v)
	})
}

func Intersect(s, t Set[T]) Set[T] {
	return MakeSet(func(v T) bool {
		return s.Contains(v) && t.Contains(v)
	})
}

func Complement(s Set[T]) Set[T] {
	return MakeSet(func(v T) bool {
		return !s.Contains(v)
	})
}

func Of(f, g func(T) bool) func(T) bool {
	return MakeFunc(Union(MakeSet(f), MakeSet(g)))
}

func And(f, g func(T) bool) func(T) bool {
	return MakeFunc(Intersetion(MakeSet(f), MakeSet(g)))
}

func Not(f func(T) bool) func(T) bool {
	return MakeFunc(Complement(MakeSet(f)))
}

type S interface {
	X()
	Y()
	Z()
}

type S interace {
	X()
}

type T interface {
	Y()
}

type U interface {
	S
	T
}

type X interface {
	X()
}
