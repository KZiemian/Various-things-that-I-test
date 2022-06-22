package main

import "fmt"

func main() {
	b := make([]byte, 32)
	read(b)
	fmt.Printf("%q\n", b)
}

func read(b []byte) {
	b[0] = 65
}
