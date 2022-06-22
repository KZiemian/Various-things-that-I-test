func Open(name string) (file *File, err error)

f, err := os.Open("filename.txt")

if err != nil {
	log.Fatal(err)
}

type error interface {
	Error() string
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func New(text string) error {
	return &errorString{text}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	//
}
