package main

import "fmt"

func main() {
	var buffer [256]byte
	var slice []byte = buffer[100:150]

	fmt.Printf("buffer: %T, %v, %v\n", buffer, len(buffer), cap(buffer))
	fmt.Printf("slice: %T, %v, %v\n", slice, len(slice), cap(slice))

	var sliceOne = buffer[100:150]
	sliceTwo := buffer[100:150]

	fmt.Printf("sliceOne: %T, %v, %v\n", sliceOne, len(sliceOne),
		cap(sliceOne))
	fmt.Printf("sliceTwo: %T, %v, %v\n", sliceTwo, len(sliceTwo),
		cap(sliceTwo))
}
