package main

import "fmt"

func main() {
	var s []int

	for i := 0; i < 20; i++ {
		fmt.Printf("len: %2d, cap: %2d\n", len(s), cap(s))

		s = append(s, i)
	}
}
