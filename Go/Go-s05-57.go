package main

import "fmt"

func main() {
	b := read()

	b[0] = 1
	fmt.Println("b:", b)
}

func read() []byte {
	b := make([]byte, 32)

	return b
}
