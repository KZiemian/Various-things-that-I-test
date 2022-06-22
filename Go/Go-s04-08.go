package main

import "fmt"

func SubtractOneFromLength(slice []byte) []byte {
	slice = slice[0 : len(slice)-1]
	return slice
}

func main() {
	slice := make([]byte, 5)


	fmt.Println("Before: len(slice) =", len(slice))
	newSlice := SubtractOneFromLength(slice)
	fmt.Println("After: len(slice) =", len(slice))
	fmt.Println("After: len(newSlice) =", len(newSlice))
}
