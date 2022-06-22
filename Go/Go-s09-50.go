type ByteReader interface {
	ReadByte() (byte, error)
}

type ByteScanner interface {
	ByteReader
	UnreadByte() error
}

type ByteWriter interface {
	WriteByte(c byte) error
}

type Closer interface {
	Close() error
}

type LimitedREader struct {
	R Reader // underlying reader
	N int64  // max bytes remaining
}

func (l *LimitedReader) Read(p []byte) (n int, err error)

func (r *PipeReader) Close() error

func (r *PipeReader) CloseWithError(err error) error

func (r *PipeReader) Read(data []byte) (n int, err error)

func (w *PipeWriter) Close() error

func (w *PipeWriter) CloseWithError(err error) error

type ReadCloser interface {
	Reader
	Closer
}

func NopCloser(r Reader) ReadCloser

type ReadSeekCloser interface {
	Reader
	Seeker
	Closer
}

type ReadSeeker interface {
	Reader
	Seeker
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

type ReadWriteSeeker interface {
	Reader
	Writer
	Seeker
}

type ReadWriter interface {
	Reader
	Writer
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type ReaderAt interface {
	ReadAt(p []byte, off int64) (n int, err error)
}

type ReaderFrom interface {
	ReadFrom(r Reader) (n int64, err error)
}

type RuneReader interface {
	ReadRune() (r rune, size int, err error)
}

type RuneScanner interface {
	RuneReader
	UnreadRune() error
}

func NewSectionReader(r Reader, off int64, n int64) *SectionReader

func (s *SectionReader) Read(p []byte) (n int, err error)

func (s *SectionReader) Seek(offset int64, whence int) (int64, error)

func (s *SectionReader) Size() int64

type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}
