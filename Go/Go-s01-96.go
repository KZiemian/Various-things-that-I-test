package main

import "fmt"

func main() {
	byteSlice := []byte("café")
	fmt.Println("byteSlice")
	fmt.Println(byteSlice)
	fmt.Printf("%v, %#v, %T\n\n", byteSlice, byteSlice, byteSlice)

	runeSlice := []rune("café")
	fmt.Println("runeSlice")
	fmt.Println(runeSlice)
	fmt.Printf("%v, %#v, %T\n", runeSlice, runeSlice, runeSlice)
}
