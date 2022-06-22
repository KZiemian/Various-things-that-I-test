package main

import "fmt"

type sliceHeader struct {
	Length        int
	ZerothElement *byte
}

func main() {
	var buffer [256]byte

	slice := sliceHeader{
		Length:       50,
		ZerothElement &buffer[100],
	}

	// slice = slice[5:10]

	slice = sliceHeader{
		Length:       48,
		ZerothElement &buffer[101],
	}

	slashPos := bytes.IndexRune(slice, '/')
}
