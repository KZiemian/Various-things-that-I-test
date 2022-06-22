package main

import "fmt"

func main() {
	var sliceByte []byte = []byte("café")
	var sliceRune []rune = []rune("café")

	fmt.Printf("sliceByte: %v, %#v, %T\n", sliceByte, sliceByte,
		sliceByte)
	fmt.Printf("sliceRune: %v, %#v, %T\n", sliceRune, sliceRune,
		sliceRune)
}
