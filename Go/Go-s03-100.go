package main

import "fmt"

func copyFirstLetters(inputString string) []byte {
	b := []byte(inputString)
	fmt.Printf("b: %v, %v, %v\n", b, len(b), cap(b))
	b = b[0:3]
	c := append([]byte{}, b...)

	return c
}

func main() {
	result := copyFirstLetters("abcdefghi")
	fmt.Printf("result: %v, %v, %v\n", result, len(result), cap(result))
}
