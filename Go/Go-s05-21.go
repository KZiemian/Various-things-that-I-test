package main

import "fmt"

type MyUInt16 uint16

const Three = 3
const TypedThree uint16 = 3

func main() {
	var mui16 MyUInt16
	mui16 = 3
	mui16 = Three
	// mui16 = TypedThree
	fmt.Println(mui16)
}
