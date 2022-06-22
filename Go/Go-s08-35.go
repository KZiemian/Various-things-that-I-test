type StringableSignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
	String() string
}

type MyInt int

func (mi MyInt) String() string {
	return fmt.Sprintf("MyInt(%d)", mi)
}

type byteseq interface {
	string | []byte
}

func Join[T byteseq](a []T, sep []T) (ret T) {
	if len(a) == 0 {
		return ret
	}
	if len(a) == 1 {
		return T(append([]byte(nil), a[0]...))
	}
	n := len(sep) * (len(a) - 1)
	for _, v := range a {
		n += len(v)
	}

	b := make([]byte, n)
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}

	return T(b)
}

func Fatal(v ...interface{})

func Fatalf(format string, v ...interface{})

func Fatalln(v ...interface{})

type structField interface {
	struct { a int; x int } |
		struct { b int; x float64 } |
		struct { c int; x uint64 }
}

func IncrementX[T structField](p *T) {
	v := p.x // INVALID
	v++
	p.x = v
}

type sliceOrMap interface {
	[]int | map[int]int
}

func Entry[T sliceOrMap](c T, i int) int {
	return c[i]
}

type sliceOrFloatMap interface {
	[]int | map[float64]int
}

func FloatEntry[T sliceOrFloatMap](c T) int {
	return c[1.0] // INVALID
}

type SliceConstraint[T any] interface {
	[]T
}

func Map[S SliceConstraint[E], E any](s S, f func(E) E) S {
	f := make(S, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

type MySlice []int

func DoubleMySlice(s MySlice) MySlice {
	v := Map[MySlice, int](s, func(e int) int { return 2 * e })
	return v
}

func SetOutput(w io.Writer)

func (l *Logger) SetOutput(w io.Writer)

type structField interface {
	struct { a int; x int } |
		struct { b int; x float64 } |
		struct { c int; x uint64 }
}

func IncrementX[T structField](p *T) {
	v := p.x // INVALID
	v++
	p.x = v
}

type sliceOrMap interface {
	[]int | map[int]int
}

type sliceOrFloatMap interface {
	[]int | map[float64]int
}

func FloatEntry[T sliceOrFloatMap](c T) int {
	return c[1.0] // INVALID
}
