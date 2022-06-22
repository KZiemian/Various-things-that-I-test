// package main

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

func Sort[Elem ?](list []Elem)

func Sort[Elem interface{ Less(y Elem) bool }](list []Elem)

type book struct {
	// ...
}

func (x book) Less(y book) bool {
	// ...
}

var bookshelf []book

Sort[book](bookshelf)

Sort[book](bookshelf)

Sort[Elem interface { Less(y Elem) bool }](list []Elem)

Sort[book interface { Less(y book) bool }](list []book)

#Sort[book]


booksort := Sort[book]

booksort(bookshelf)

type Lesser[T any] interface {
	Less(y T) bool
}

func Sort[Elem Lesser[Elem]](list []Elem)

func Sort[Elem interface { Less(y Elem) bool }](list []Elem) {
	// ...
	var i, j int
	// ...
	if list[i].Less(list[j]) {
		// ...
	}
	// ...
}

Sort([]int{1, 2, 3})

type myInt int

func (x myInt) Less(y myInt) bool {
	return x < y
}

type Float interface {
	type float32, float64
}

func Sin[T Float](x T) T

func min[T Ordered](x, y T) T ...

type Ordered interface {
	type int, int8, int16, ..., uint, uint8, uint16, ...,
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
		// ...
		}
	// ...
}

type Bytes interface {
	type []byte, string
}

func Index[bytes Bytes](s, sep bytes) int

type Pointer[T any] interface {
	type *T
}

func f[T any, PT Pointer[T]](x T)

func foo[T any, PT interface {type *T}](x T)


func BasicSort[Elem Ordered](list []Elem)

func Sort[Elem Lesser[Elem]](list []Elem)

type Lesser[Elem any] interface {
	Less(Elem) Elem
}

func ReadAll(r io.Reader)(b []byte, err error)

func ReadAll[reader io.Reader](r reader) (b []byte, err error)

func Drain[T any](c <-chan T)

func Merge[T any](c1, c2 <-chan T) <-chan T
