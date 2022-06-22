package main

import "fmt"

func main() {
	s := []int{1, 1, 1}

	for i := range s {
		s[i] += 1
	}

	fmt.Println("s:", s)
}
