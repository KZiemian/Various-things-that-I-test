package main

import "fmt"

func main() {
	b := make([]byte, 32)
	read(b)
	fmt.Println("b:", b)
}

func read(b []byte) {
	b[0] = 1
}
