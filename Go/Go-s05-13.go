package main

import "fmt"

const Two = 2.0 + 0i

func main() {
	s := Two
	fmt.Printf("%T: %v\n", s, s)

	var f float64
	var g float64 = Two
	f = Two
	fmt.Println(f, "and", g)
}
