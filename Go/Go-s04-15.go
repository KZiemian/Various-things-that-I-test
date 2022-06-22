package main

import "fmt"

type sliceHeader struct {
	Length        int
	Capacity      int
	ZerothElement *byte
}

func main() {
	var iBuffer [10]int
	slice := iBuffer[0:0]

	slice := sliceHeader{
		Length:        0,
		Capacity:      10,
		ZerothElement: &iBuffer[0],
	}
}
