err := errors.New("kerbroom")
fmt.Printf("%v\n", err)

err := errors.Wrap(
	syscall.EBADF, "couldn't write to stream")
fmt.Printf("%v\n", err)

fmt.Printf("%+v\n", err)

func Errorf(format string, args ...interface{}) error

func Wrapf(err error, format string, args ...interface{}) error

func Cause(err error) error

type temporary interface {
	Temporary() bool
}

func IsTemporary(err error) bool {
	te, ok := errors.Cause(err).(temporary)
	return ok && te.Temporary()
}

func Write(w io.Writer, buf []byte) error {
	_, err := w.Write(buf)
	if err != nil {
		log.Println("unable to write:", err)

		return err
	}
	return nil
}

func Write(w io.Write, buf []byte) error {
	_, err := w.Write(buf)
	return errors.Wrap(err, "write failed")
}
