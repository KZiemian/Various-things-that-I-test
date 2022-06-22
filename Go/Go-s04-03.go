package main

import "fmt"

func main() {
	var buffer [512]byte

	fmt.Printf("buffer: %T, %v, %v\n", buffer, len(buffer), cap(buffer))
}
