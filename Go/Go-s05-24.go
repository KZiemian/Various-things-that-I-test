package main

import "fmt"

type MyUintptr uintptr

const Three = 3
const TypedThree uintptr = 3

func main() {
	var muip MyUintptr
	muip = 3
	muip = Three
	// muip = TypedThree
	fmt.Println(muip)
}
