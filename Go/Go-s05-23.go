package main

import "fmt"

type MyUInt64 uint64

const Three = 3
const TypedThree uint64 = 3

func main() {
	var mui64 MyUInt64
	mui64 = 3
	mui64 = Three
	// mui64 = TypedThree
	fmt.Println(mui64)
}
