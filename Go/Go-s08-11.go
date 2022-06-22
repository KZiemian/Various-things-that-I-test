func Print(s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

Print[int]([]int{1, 2, 3})

// This function is INVALID.
func Stringify[T any](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String()) // INVALID
	}
	return ret
}

type Stringer interface {
	String() string
}

func Print[T interface{}](s []T) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ren
}

func Stringify[T Stringer](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}

func Print2[T1, T2 any](s1 []T1, s2 []T2) {
	// ...
}

func Print2Same[T any](s1 []T, s2 []T) {
	// ...
}

func Stringer interface {
	String() string
}

type Plusser interface {
	Plus(string) string
}

func ConcatTo[S Stringer, P Plusser](s []S, p []P) []string {
	r := make([]string, len(s))
	for i, v := range s {
		r[i] = p[i].Plus(v.String())
	}
	return r
}

func Stringify2[T1, T2 Stringer](s1 []T1, s2 []T2) string {
	r := ""
	for _, v1 := range s1 {
		r += v1.String()
	}
	for _, v2 := range s2 {
		r += v2.String()
	}
	return r
}

type Vector[T any] []T

var v Vector[int]

func (v *Vector[T]) Push(x T) {
	*v = append(*v, x)
}

type List[T any] struct {
	next *List[T]
	val  T
}

// This type is INVALID.
type P[T1, T2 any] struct {
	F *P[T2, T1]		// INVALID; must be [T1, T2]
}

type ListHead[T any] struct {
	head *ListElement[T]
}

type ListElement[T any] struct {
	next *ListElement[T]
	val  T
	head *ListHead[T]
}

type StringableVector[T Stringer] []T

func (s StringableVector[T]) String() string {
	var sb strings.Builder
	for i, v := range s {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(v.String())
	}
	return sb.String()
}

// This fuction is INVALID.
func Smallest[T any](s []T) T {
	r := s[0] // panic if slice is empty
	for _, v := range s[1:] {
		if v < r { 	// INVALID
			r = v
		}
	}
	return r
}
