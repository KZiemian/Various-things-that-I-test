func TestIndex(t *testing.T) {
	var tests = []struct {
		s   string
		sep string
		out int
	}{
		{"", "", 0},
		{"", "a", -1},
		{"fo", "foo", -1},
		{"foo", "foo", 0},
		{"oofofoofooo", "f", 2},
		// etc.
	}

	for _, test := range tests {
		actual := strings.Index(test.s, test.sep)
		if actual != test.out {
			t.Errorf("Index(%q, %q) = %v; want %v", test.s,
				test.sep, actual, test.out)
		}
	}
}

var (
	viewCount   int64
	viewCountMu sync.Mutex
)

var viewCount struct {
	sync.Mutex
	n int64
}

viewCount.Lock()
viewCount.n++
viewCount.Unlock()

return struct {
	io.ReadSeeker
	io.Closer
}{
	io.NewSectionReader(strings.NewReader(s), 0, int64(len(s))),
	ioutil.NopCloser(nil),
}

var s interface {
	String() string
} = bytes.NewBufferString("I'm secretly a fmt.Stringer!")
fmt.Println(s.String())

if _, ok := stdin.(interface {
	Fd() uintptr
}); !ok {
	t.Error("can't access methods of underlying *os.File")
}

var f func(*bytes.Buffer, string) (int, error)
var buf bytes.Buffer
f = (*bytes.Buffer).WriteString
f(&buf, "y u no buf.WriteString?")
buf.WriteTo(os.Stdout)

var f func(string) (int, error)
var buf bytes.Buffer
f = buf.WritesString
f("Hey... ")
f("this *is* cute.")
buf.WriteTo(os.Stdout)



func Double[E contstraints.Number](s []E) []E {
	r := make([]E, len(s))
	for i, v := range s {
		r[i] = v + v
	}
	return r
}

type MySlice []int

var V1 = Double(MySlice{1})

type SC[E any] interface {
	[]E
}

func DoubleDefined[S SC[E], E constraints.Number](s S) S {
	r := make(S, len(s))
	for i, v := range s {
		r[i] = v + v
	}
	return r
}

var V2 = DoubleDefined[MySlice, int](MySlice{1})

var V3 = DoubleDefined(MySlice{1})

type Setter interface {
	Set(string)
}

func FromStrings[T Setter](s []string) []T {
	result := make([]T, len(s))
	for i, v := range s {
		result[i].Set(v)
	}
	return result
}

type Settable int

func (p *Settable) Set(s string) {
	i, _ := strconv.Atoi(s)  // real code should not ignore the error
	*p = Settable(i)
}

func F() {
	// INVALID
	nums := FromStrings[Settable]([]string{"1", "2"})
	// Here we want nums to be []Settable{1, 2}.
	// ...
}

func F() {
	// Compiles but does not work ad desired.
	// This will panic at run time when calling the Set method.
	nums := FromStrings[*Settable]([]string{"1", "2"})
	// ...
}

type Setter2[B any] interface {
	Set(string)
	*B // non-interface type constraint element
}

func FromStrings2[T any, PT Setter2[T]](s []string) []T {
	result := make([]T, len(s))
	for i, v := range s {
		p := PT(&result[i])
		p.Set(v)
	}
	return result
}

func F2() {
	nums := FromStrings2[Settable, *Settable]([]string{"1", "2"})
	// ...
}

func F3() {
	nums := FromStrings2[Settable]([]string{"1", "2"})
	// ...
}
