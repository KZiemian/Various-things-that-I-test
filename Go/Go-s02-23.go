package main

import "fmt"

func main() {
	var s []int

	for i := 0; i < 10; i++ {
		fmt.Printf("len: %2d, cap: %2d, %v\n", len(s), cap(s), s)

		s = append(s, i)
	}
}
