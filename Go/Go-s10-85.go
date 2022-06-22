// When to use generics?
// 1. Imporved static type safety.
// 2. More efficient memory use.
// 3. (Significantly) better performance.

[P, Q constraint_1, R contraint_2]

func min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

m := min[int](2, 3)

fmin := min[float64]

m := fmin(2.71, 3.14)

type Tree[T interface{}] struct {
	left, right *Tree[T]
	data        T
}

func (t *Tree[T]) Lookup(x T) *Tree[T]

var stringTree Tree[string]

func min[T constraints.Ordered](x, y T) T

func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

interface {
	int|string|bool
}

package constraints

type Ordered interface {
	Integer|Float|~string
}

[S interface{~[]E}, E interface{}]

[S ~[]E, E interface{}]

[S ~[]E, E any]

func min[T constraints.Ordered](x, y T) T

var a, b, m float64

m = min[float64](a, b)

m = min(a, b)

func Scale[E constraints.Integer](s []E, c E) []E {
	r := make([]E, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

type Point []int32

func (p Point) String() string {
	// ...
}

func ScaleAndPrint(p Point) {
	r := Scale(p, 2)
	fmt.Println(r.String()) // DOES NOT COMPILE
}

func Scale[S ~[]E, E constraints.Integer](s S, sc E) S {
	r := make(S, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

func Scale[S ~[]E, E constraints.Integer](s S, sc E) S {
	r := make(S, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}
