package main

import "fmt"

func main() {
	s := []int{1, 1, 1}

	for _, n := range s {
		n += 1
	}

	fmt.Println("s:", s)
}
