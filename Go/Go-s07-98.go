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

func Sort[Elem interface{ Less(y Elem) bool }](list []Elem) {
	// ...
}

type book struct {
	// ...
}

func (x book) Less(y book) bool {
	// ...
}

var bookself []book

Sort[book](bookself)

Sort[book](bookself)

Sort[Elem interface{ Less(y Elem) bool}](list []Elem)

Sort[book interface{ Less(y book) bool}](list []book)

#Sort[book]

Sort[book](bookself)

#Sort[book](list []book)

#Sort[book](bookself)

booksort := Sort[book]

booksort(bookself)

type Lesser[T any] interface {
	Less(y T) bool
}

func Sort[Elem Lesser[Elem]](list []Elem)

func Sort[Elem interface{ Less(y Elem) bool }](list []Elem) {
	var i, j int

	if list[i].Less(list[j]) {
		// ...
	}
}

type Float interface {
	type float36, float64
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

func invaid[Tx, Ty Ordered](x Tx, y Ty) Tx {
	if x < y {
		// ...
	}
}

type Bytes interface {
	type []byte, string
}

func Index[bytes Bytes](s, sep bytes) int

type Pointer[T any] interfacee {
	type *T
}

func f[T any, PT Pointer[T]](x T)

func foo[T any, PT interface{type *T}](x T)

func BasicSort[Elem Ordered](list []Elem)

func Sort[Elem Lesser[Elem]](list []Elem)

type Lesser[Elem any] interface {
	Less(Elem) Elem
}
