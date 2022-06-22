package main

import "fmt"

type MyInt16 int16

const Three = 3
const TypedThree int16 = 3

func main() {
	var mi16 MyInt16
	mi16 = 3
	mi16 = Three
	// mi16 = TypedThree
	fmt.Println(mi16)
}
