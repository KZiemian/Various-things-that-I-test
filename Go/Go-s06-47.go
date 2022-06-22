package main

func main() {
	if err != nil {
		return ...
	}

	if err != nil {
		return annotate(err)
	}

	fd, err := sql.OpenDB(conn)
	if err != nil {
		return nil, err
	}

	fd := check sql.OpenDB(conn)

	handle err {
		log.Fatal(err)
	}

	handle err {
		return fmt.Errorf("couldn't open database: %v", err)
	}
}

type Wrapper interface {
	Unwrap() error
}

func Is(err, target error) bool

func As[type E](err error) (e E, ok bool)

type Formatter interface {
	Format(p Printer) (next error)
}

func Uniq(type T)(<-chan T) <-chan T

func Print(type T)(s []T) {
	for _, v := range s {
		fmt.Println()
	}
}

stream := make(chan bool)
filtered := Uniq(bool)(stream)
Print(int)([]int{1, 2, 3})

func Min(type T)(a, b T) T {
	if a < b {
		return a
	}
	return b
}

contract Less(t T) {
	t < t
}

func Min(type T Less(T))(a, b T) T {
	if a < b {
		return a
	}
	return b
}

m := Min[int](3.0, 4)

type Vector(type Element) []Element
var v Vector(int)

func (v *Vector(T)) Push(x T) {
	*v = append(*v, x)
}
