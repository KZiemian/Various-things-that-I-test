var bookshelf []book

Sort[book](bookshelf)

Sort[book]

Sort[Elem interface{ Less(y Elem) bool }](list []Elem)

Sort[book interface{ Less(y Elem) bool }](list []Elem)

#Sort[book]

booksort := Sort[book]
booksort(bookshelf)

type Lesser[T any] interface {
	Less(y T) bool
}

func Sort[Elem Lesser[Elem]](list []Elem)

func Sort[Elem interface{ Less(y Elem) bool }](list []Elem) {
	// ...
	var i, j int
	// ...
	if list[i].Less(list[j]) {
		// ...
	}
	// ...
}

Sort([]int{1, 2, 3})

type myInt interface

func (x myInt) Less(y myInt) bool {
	return x < y
}

type Float interface {
	type float32, float64
}

func Sin[T Float](x T) T

func min[T Ordered](x, y T) T {
	// ...
}

type Ordered interface {
	type int, int8, int16, ..., uint8, uint16, ...,
		float32, float64, string
}

func min[T Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func invalid[Tx, Ty Ordered](x Tx, y Ty) Tx {
	// ...
	if x < y {
		// INVALID
	}
	// ...
}

type Bytes interface {
	type []byte, string
}

func Index[bytes Bytes](s, sep bytes) int {
	// ...
}

type Pointer[T any] interface {
	type *T
}

func f[T any, PT Pointer[T]](x T)

func foo[T any, PT interface{type *T}](x T)

func Basic[Elem Ordered](list []Elem)

func Sort[Elem Lesser[Elem]](list []Elem)

type Lesser[Elem any] interface {
	Less(Elem) Elem
}

func ReadAll(r io.Reader) ([]byte, error)

func ReadAll[reader io.Reader](r reader) ([]byte, error)

func Drain[T any](c <-chan T)

func Merge[T any](c1, c2 <-chan T) <-chan T
