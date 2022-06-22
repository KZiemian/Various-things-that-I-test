func ComputeCost(X, y, theta mat.Matrix) float64 {
	h := mat.Minus(mat.Product(X, theta), y)

	return mat.Dot(h, h).Sum() / float64(2 * X.Rows())
}

if err != nil {
	return fmt.Errorf("whoops: %v", err)
}

return &PathError{"open", filename, err}

p, ok := err.(*PathError)

if p.Err == ... {
	// ...
}

if os.IsPremission(err) {
	// ...
}

package errors

type Wrapper interface {
	Unwrap() error
}

func (e *PathError) Unwrap() error {
	return e.Err
}

func Is(err, target error) bool

if errors.Is(err, io.EOF) {
	// ...
}

if errors.Is(err, os.ErrPermision) {
	// ...
}

if os.IsPermission(err) {
	// ...
}

if errors.Is(err, net.ErrTimeout) {
	// ...
}

if err, ok := err.(net.Error); ok && err.Timeout() {
	// ...
}

func As(err error, target interface{}) bool

var p *os.PathError
if errors.As(err, &p) {
	// ...
}

if p := (*os.PathError)(nil); errors.As(err, &p) {
	// ...
}

fmt.Printf("%v", err) {
	// ...
}

fmt.Printf("%+v", err) {
	// ...
}

func ReverseInts(s []int) {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}

func ReverseStrings(s []string) {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}

func Reverse[type Element](s []Element) {
	first := 0
	last := len(s) - 1

	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}

func ReverseAndPrint(s []int) {
	fmt.Println(Reverse[int](s))
}
