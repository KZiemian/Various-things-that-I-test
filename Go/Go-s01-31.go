package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("s: %v, %#v, %T\n", s, s, s)

	s = s[1:4]
	fmt.Printf("s: %v, %#v, %T\n", s, s, s)

	s = s[:2]
	fmt.Printf("s: %v, %#v, %T\n", s, s, s)

	s = s[1:]
	fmt.Printf("s: %v, %#v, %T\n", s, s, s)
}
