package main

import "fmt"

type MyInt int

const Three = 3
const TypedThree int = 3

func main() {
	var mi MyInt
	mi = 3
	mi = Three
	// mi = TypedThree
	fmt.Println(mi)
}
