package some

var digitRegexp = regexp.MustCompile("[0-9]+")

func CopyDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)

	return c
}
