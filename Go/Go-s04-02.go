package main

import "fmt"

func main() {
	var buffer [256]byte

	fmt.Printf("buffer: %v\n", buffer)
	fmt.Printf("buffer: %T, %v, %v\n", buffer, len(buffer), cap(buffer))
}
