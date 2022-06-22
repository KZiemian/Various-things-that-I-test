package main

import "fmt"

type MyUInt8 uint8

const Three = 3
const TypedThree uint8 = 3

func main() {
	var mui8 MyUInt8
	mui8 = 3
	mui8 = Three
	// mui8 = TypedThree
	fmt.Println(mui8)
}
