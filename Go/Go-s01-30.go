package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11}
	fmt.Printf("q: %v, %#v, %T\n\n", q, q, q)

	r := []bool{true, false, true, false, true}
	fmt.Printf("r: %v, %#v, %T\n\n", r, r, r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Printf("s: %v, %#v, %T\n", s, s, s)
}
