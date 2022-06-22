package main

import "fmt"

func main() {
	s := fmt.Sprintf("%[2]d %[1]d", 11, 22)
	fmt.Printf("s:%s\n", s)

	s = fmt.Sprintf("%[3]*.[2]*[1]f", 12.0, 2, 6)
	fmt.Printf("s:%s\n", s)

	s = fmt.Sprintf("%6.2f", 12.0)
	fmt.Printf("s:%s\n", s)

	s = fmt.Sprintf("%d %d %#[1]x %#x", 16, 17)
	fmt.Printf("s:%s\n", s)
}
