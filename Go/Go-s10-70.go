type StringWriter interface {
	WriteString(s string) (n int, err error)
}

type WriteCloser interface {
	Writer
	Closer
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

var Discard Writer = discard{}

type WriteAt interface {
	WriteAt(p []byte, off int64) (n int, err error)
}

type WriterTo interface {
	WriterTo(w Writer) (n int64, err error)
}

type X string

func (x X) String() string {
	return Sprintf("<%s>", string(x))
}

func Fscan(r io.Reader, a ...interface{}) (n int, err error)

type Formatter interface {
	Format(f State, verb rune)
}

type ScanState interface {
	ReadRune() (r rune, size int, err error)
	UnreadRune() error
}

SkipSpace()

Token(skipSpace bool, f func(rune) bool) (token []byte, err error)

Width() (wid int, ok bool)

Read(buf []byte) (n int, err error)

type Scanner interface {
	Scan(state ScanState, verb rune) error
}

type State interface {
	Write(b []byte) (n int, err error)
	Width() (wid int, ok bool)
	Precision() (prec int, ok bool)
	Flag(c int) bool
}



select {
case v1 := <-c1:
	fmt.Printf("received %v from c1\n", v1)
case v2 := <-c2:
	fmt.Printf("received %v from c2\n", v2)
case c3 <- 23:
	fmt.Printf("sent %v to c3\n", 23)
default:
	fmt.Printf("no one was ready to communicate\n")
}


func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}


for i := 0; i < n; i++ {
	for j := 0; j < p; j++ {
		var t float64
		for k := 0; k < m; k++ {
			t += a.AT__(i, k) *
				b.AT__(k, j)
		}
		c.ATSET__(i, j, t)
	}
}

map[KeyType]ValueType
