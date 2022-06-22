package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(b []byte) (n int, err error) {
	text := make([]byte, 8)
	n, err = rot13.r.Read(text)
	if err != nil {
		return n, nil
	}

	for i, elem := range text {
		if (65 <= elem) && (elem <= 77) {
			b[i] = elem + 13
		} else if (78 <= elem) && (elem <= 90) {
			b[i] = 65 + ((elem - 78) % 13)
		} else if (97 <= elem) && (elem <= 109) {
			b[i] = elem + 13
		} else if (110 <= elem) && (elem <= 122) {
			b[i] = 97 + ((elem - 97) % 13)
		} else {
			b[i] = elem
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
