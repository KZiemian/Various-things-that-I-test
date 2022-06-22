func (e *SyscallError) Error() string

func (e *SyscallError) Timeout() bool

func (e *SyscallError) Unwrap() error

var EOF = errors.New("EOF")

var ErrClosedPipe = errors.New("io: read/write on closed pipe")

var ErrNoProgress = errors.New("multiple Read calls return no data or error")

var ErrShortBuffer = errors.New("short buffer")

var ErrShortWrite = errors.New("short write")

var ErrUnexpectedEOF = errors.New("unexpected EOF")

type ReadWriterCloser interface {
	io.ReadCloser
	io.WriteCloser
}

type MyCloser interface {
	Close()
}

type MyReadCloser interface {
	io.ReadCloser
	MyCloser
}

func main() {
	err := readConfig()
	if err == os.ErrNotExist {
		// ...
	}
}

func readConfig() error {
	f, err := os.Open(`config json`)
	if err != nil {
		return fmt.Errorf(`Cannot open config json: %w`, err)
	}
}

return fmt.Errorf("Call failed: %v", err)
fmt.Printf("%T", err) // *err.stringError

return fmt.Errorf("Call failed: %w", err)
fmt.Printf("%T", err) // *err.wrapError

err := readConfig()
if errors.Is(err, os.ErrNotExist) {
	// file does not exist
}

err := readDatabase()
var e *QueryError
if errors.As(err, &e) {
	// err is type QueryError
	// also works with wrap!
}

reflect.ValueOf(0).IsZero() // ture
reflect.ValueOf("").IsZero() // true
reflect.ValueOf(struct{}{}).IsZero() // true

reflect.ValueOf([]string{}).IsZero() // false
reflect.ValueOf([]string(nil)).IsZero() // true

func Test_something(t *testing.T) {
	t.CleanUp(func() {
		fmt.Println("Clean up!")
	})

	t.Run(tt.name, [...])
}

import "github.com/quux/bar"

func fn() error {
	x, err := bar.Foo()
	if err != nil {
		return err
	}
	// use x
}

type temporary interface {
	Temporary() bool
}

func IsTemporary(err error) bool {
	te, ok := err.(temporary)
	return ok && te.Temporary()
}

func AuthenticateRequest(r *Request) error {
	err := authenticate(r.User)
	if err != nil {
		return fmt.Errorf("authenticate failed: %v", err)
	}
	return nil
}
