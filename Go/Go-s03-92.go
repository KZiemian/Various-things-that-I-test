package main

import "fmt"

func main() {
	s := []byte{'g', 'o', 'l', 'a', 'n', 'g'}

	fmt.Printf("s: %v, %v, %v\n", s, len(s), cap(s))

	t := make([]byte, len(s), (cap(s)+1)*2)

	fmt.Printf("t: %v, %v, %v\n", t, len(t), cap(t))
	copy(t, s)
	s = t
	fmt.Printf("t: %v, %v, %v\n", t, len(t), cap(t))
	fmt.Printf("s: %v, %v, %v\n", s, len(s), cap(s))
}
