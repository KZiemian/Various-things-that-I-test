func MapKeys(m map[string]int) []string {
	var s []string
	for k := range m {
		s = append(s, k)
	}
	return s
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	var s []K
	for k := range m {
		s = append(s, k)
	}
	return s
}

type Tree[T any] struct {
	cmp  func(T, T) int
	root *leaf[T]
}

type leaf[T any] struct {
	val         T
	left, right *left[T]
}

func (bt *Tree[T]) find(val T) **leaf[T] {
	pl := &bt.root
	for *pl != nil {
		switch cmp := bt.cmp(val, (*pl).val); {
		case cmp < 0:
			pl = &(*pl).left
		case cmp > 0:
			pl = &(*pl).right
		default:
			return pl
		}
	}
	return pl
}

type SliceFn[T any] struct {
	s   []T
	cmp func(T, T) bool
}

func (s SliceFn[T]) Len() int {
	return len(s.s)
}

func (s SliceFn[T]) Swap(i, j int) {
	s.s[i], s.s[j] = s.s[j], s.s[i]
}

func (s SliceFn[T]) Less(i, j int) bool {
	return s.cmp(s.s[i], s.s[j])
}

func SortFn[T any](s []T, cmp func(T, T) bool) {
	sort.Sort(SliceFn[T]{s, cmp})
}

func ReadFour[T io.Reader](r T) ([]byte, error) {
	buf := make([]byte, 4)
	_, err := io.ReadFull(r, buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

a
_x9
ThisVariableIsExported
\alpha\beta

imaginary_lit = (decimal_digits | int_lit | float_lit) "i" .
\a \b \f \n \r \t \v \\ \' \"

rune_lit = "'" ( unicode_value | byte_value ) "'" .
unicode_value = unicode_char | little_u_value | big_u_value | escaped_char .
byte_value = octal_byte_value | hex_byte_value .
octal_byte_value = `\` octal_digit octal_digit octal_digit .
hex_byte_value = `\` "x" hex_digit hex_digit.
little_u_value = `\` "u" hex_digit hex_digit hex_digit hex_digit .
big_u_value = `\` "U" hex_digit hex_digit hex_digit hex_digit
hex_digit hex_digit hex_digit hex_digit .
escaped_char = `\` ( "a" | "b" | "f" | "n" | "r" | "t" | "v" | `\` | "'" | `"` ) .

if err != nil {
	return fmt.Errorf("whoophs: %v", err)
}

return &PathError{"open", filename, err}

p, ok := err.(*PathError)
if p.Err == ... {...}
if os.IsPermission(err) {...}

func ComputeCost(X, y, theta *mat.Dense) float64 {
	m, _ := X.Dims()
	h := new(mat.Dense)
	h.Mul(X, theta)
	h.Sub(h, y)
	h.MulElem(h, h)

	return mat.Sum(h) / float64(2*m)
}
