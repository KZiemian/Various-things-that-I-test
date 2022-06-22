package main

import "fmt"

func main() {
	var s []int
	fmt.Printf("s: %v, %#v, %T\n", s, s, s)

	fmt.Printf("nil: %v, %#v, %T\n", nil, nil, nil)
}
