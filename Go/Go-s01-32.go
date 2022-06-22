package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("s: %v, %#v, %T\n", s, s, s)
	fmt.Printf("len(s) = %d, cap(s) = %d, %v\n\n",
		len(s), cap(s), s)

	s = s[:0]
	fmt.Printf("s: %v, %#v, %T\n", s, s, s)
	fmt.Printf("len(s) = %d, cap(s) = %d, %v\n\n",
		len(s), cap(s), s)

	s = s[:4]
	fmt.Printf("s: %v, %#v, %T\n", s, s, s)
	fmt.Printf("len(s) = %d, cap(s) = %d, %v\n\n",
		len(s), cap(s), s)

	s = s[2:]
	fmt.Printf("s: %v, %#v, %T\n", s, s, s)
	fmt.Printf("len(s) = %d, cap(s) = %d, %v\n",
		len(s), cap(s), s)

	s = s[:5]
	fmt.Println(s)
}
