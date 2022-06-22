package main

import "fmt"

type sliceHeader struct {
	Length        int
	Capacity      int
	ZerothElement *byte
}

func main() {
	slice1Header := sliceHeader{
		Length:        0,
		Capacity:      0,
		ZerothElement: nil,
	}

	slice2Header := sliceHeader{}
}
