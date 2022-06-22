package main

import "fmt"

type sliceHeader struct {
	Length        int
	ZerothElement *byte
}

func main() {
	var buffer [256]byte

	slice := sliceHeader{
		Length:        50,
		ZerothElement: &buffer[100],
	}

	fmt.Printf("slice: %v, %T\n", slice, slice)
	fmt.Printf("slice.Length: %v, %T\n", slice.Length, slice.Length)
	fmt.Printf("slice.ZerothElement: %v, %T\n", slice.ZerothElement,
		slice.ZerothElement)


	// slice2 := slice[5:10]

	// fmt.Printf("slice2: %v, %T\n", slice2, slice2)

	slice2Header := sliceHeader{
		Length:        5,
		ZerothElement: &buffer[105],
	}

	fmt.Printf("slice2Header: %v, %T\n", slice2Header, slice2Header)
	fmt.Printf("slice2Header.Length: %v, %T\n", slice2Header.Length,
		slice2Header.ZerothElement)
}
