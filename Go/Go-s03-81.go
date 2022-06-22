package main

import "fmt"

func main() {
	var s []int

	fmt.Printf("s: %v, %#v, %T\n", s, s, s)
	fmt.Printf("s: %v, %v, %v\n", s, len(s), cap(s))

	var sOne = make([]int, 0)

	fmt.Printf("sOne: %v, %#v, %T\n", sOne, sOne, sOne)
	fmt.Printf("sOne: %v, %v, %v\n", sOne, len(sOne), cap(sOne))
}
