package main

import "fmt"

func main() {
	var (
		i int
		s string
	)

	fmt.Sscanf(" 1234567 ", "%5s%d", &s, &i)
	fmt.Printf("s: %s, i: %d\n", s, i)

	fmt.Sscanf(" 12 34 567 ", "%5s%d", &s, &i)
	fmt.Printf("s: %s, i: %d\n", s, i)
}
