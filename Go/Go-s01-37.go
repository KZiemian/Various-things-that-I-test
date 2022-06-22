package main

import "fmt"

func main() {
	var s []int
	fmt.Printf("s: %v, %#v, %T\n", s, s, s)
	fmt.Printf("len(s) = %d, cap(s) = %d, %v\n\n",
		len(s), cap(s), s)

	// append works on nil slices
	s = append(s, 0)
	fmt.Printf("s: %v, %#v, %T\n", s, s, s)
	fmt.Printf("len(s) = %d, cap(s) = %d, %v\n\n",
		len(s), cap(s), s)

	// The slice grows as needed.
	s = append(s, 1)
	fmt.Printf("s: %v, %#v, %T\n", s, s, s)
	fmt.Printf("len(s) = %d, cap(s) = %d, %v\n\n",
		len(s), cap(s), s)

	s = append(s, 2, 3, 4)
	fmt.Printf("s: %v, %#v, %T\n", s, s, s)
	fmt.Printf("len(s) = %d, cap(s) = %d, %v\n",
		len(s), cap(s), s)
}
