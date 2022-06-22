package main

import "fmt"

type MyInt32 int32

const Three = 3
const TypedThree int32 = 3

func main() {
	var mi32 MyInt32
	mi32 = 3
	mi32 = Three
	// mi32 = TypedThree
	fmt.Println(mi32)
}
