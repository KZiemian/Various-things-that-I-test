package main

import "fmt"

func main() {
	string1 := fmt.Sprintf("%[2]d %[1]d\n", 11, 22)
	fmt.Printf("string1: %v", string1)

	string1 = fmt.Sprintf("%[3]*.[2]*[1]f", 12.0, 2, 6)
	fmt.Printf("string1: %v\n", string1)

	string1 = fmt.Sprintf("%6.2f", 12.0)
	fmt.Printf("string1: %v\n", string1)

	string1 = fmt.Sprintf("%d %d %#[1]x %#x", 16, 17)
	fmt.Printf("string1: %v\n", string1)
}
