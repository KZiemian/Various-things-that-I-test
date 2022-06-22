package main

import "fmt"

type MyUInt uint

const Three = 3
const TypedThree uint = 3

func main() {
	var mui MyUInt
	mui = 3
	mui = Three
	// mui = TypedThree
	fmt.Println(mui)
}
