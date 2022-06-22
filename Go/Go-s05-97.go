type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

err := something()

switch err := err.(type) {
case nil:
	//
case *os.PathError:
	//
default:
}

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
	err := authentication(r.User)
	if err != nil {
		return fmt.Errorf("authentication failed: %v", err)
	}
	return nil
}
