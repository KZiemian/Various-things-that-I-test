package main

import "fmt"

type MyInt8 int8

const Three = 3
const TypedThree int8 = 3

func main() {
	var mi8 MyInt8
	mi8 = 3
	mi8 = Three
	// mi8 - TypedThree
	fmt.Println(mi8)
}
