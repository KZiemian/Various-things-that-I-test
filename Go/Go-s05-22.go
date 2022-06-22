package main

import "fmt"

type MyUInt32 uint32

const Three = 3
const TypedThree uint32 = 3

func main() {
	var mui32 MyUInt32
	mui32 = 3
	mui32 = Three
	// mui32 = TypedThree
	fmt.Println(mui32)
}
