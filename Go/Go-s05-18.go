package main

import "fmt"

type MyInt64 int64

const Three = 3
const TypedThree int64 = 3

func main() {
	var mi64 MyInt64
	mi64 = 3
	mi64 = Three
	// mi64 = TypdThree
	fmt.Println(mi64)
}
