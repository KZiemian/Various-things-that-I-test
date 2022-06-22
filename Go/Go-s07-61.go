var s = []string{
	i int
	s string
}{
	struct {
		i int
		s string
	}{6 * 9, "Question"},
	struct {
		i int
		s string
	}{42, "Answer"},
}

var t []struct {
	i int
	s string
}{
	{6 * 9, "Question"},
	{42, "Answer"}
}

return struct{
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

if _, ok := stding.(interface {
	Fd() uintptr
}); !ok {
	t.Error("can't access methods of underlying *os.File")
}
