package main

import "fmt"

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) {
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

func main() {
	p := []byte{2, 3, 5}

	fmt.Printf("p: %v, %v, %v, %T\n", p, len(p), cap(p), p)

	p = AppendByte(p, 7, 11, 13)

	fmt.Printf("p: %v, %v, %v, %T\n", p, len(p), cap(p), p)
}
