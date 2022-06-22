package main

import "fmt"

func copyFirstLetters(inputString string) []byte {
	b := []byte(inputString)
	b = b[0:3]
	c := make([]byte, len(b))
	copy(c, b)

	return c
}

func main() {
	result := copyFirstLetters("abcdefg")
	fmt.Printf("result: %v, %v, %v\n", result, len(result), cap(result))
}
