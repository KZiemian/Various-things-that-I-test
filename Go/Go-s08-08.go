package main

import "fmt"

func main() {
	b := read()
	fmt.Println("b:", b)
}

func read() []byte {
	b := make([]byte, 32)
	return b
}
