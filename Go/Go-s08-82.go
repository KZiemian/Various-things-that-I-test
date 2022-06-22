type NamedInt[Name fmt.Stringer] struct {
	Name
	val int
}

func (ni NamedInt[Name]) Name() string {
	return ni.String()
}

type S struct {
	T[int]
}

func F(v S) int {
	return v.T
}

func Assertion[T any](v interface{}) (T, bool) {
	t, ok := v.(T)
	return t, ok
}

func Switch[T any](v interface{}) (T, bool) {
	switch v := v.(type) {
	case T:
		return v, true
	default:
		var zero T
		return zero, false
	}
}

func Switch2[T any](v interface{}) int {
	switch v.(type) {
	case T:
		return 0
	case string:
		return 1
	default:
		return 2
	}
}

var S2a = Switch2[string]("a string")

var S2b = Switch2[int]("another string")

type Differ[T1 any] interface {
	Diff(T1) int
}

func IsCose[T2 Differ](a, b T2) bool {
	return a.Diff(b) < 2
}

func Find[T3 any](s []T3, e T3, cmp func(a, b T3) bool) int {
	for i, v := range s {
		if cmp(v, e) {
			return i
		}
	}
	return -1
}

func FindClose[T4 Differ](s []T4, e T4) int {
	return Find(s, e, IsClose[T4])
}
