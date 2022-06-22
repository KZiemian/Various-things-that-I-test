type Reader interface {
	Read(buf []byte) (n int, err error)
}

func Save(f *os.File, doc *Document) error

func Save(rwc io.ReadWriteCloser, doc *Document) error

func Save(wc io.WriteCloser, doc *Document) error

type NopCloser struct {
	io.Writer
}

func (c *NopCloser) Closer() error {
	return nil
}

func Save(w io.Writer, doc *Document) error
