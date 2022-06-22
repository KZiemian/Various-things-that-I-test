package main

import "fmt"

type MyBool bool
const True = true
const TypedTrue bool = True

func main() {
	var mb MyBool
	mb = true
	mb = True
	// mb = TypedTrue
	fmt.Println(mb)
}
