package main

import "fmt"

const hello = "Hello, 世界"

type MyString string

func main() {
	fmt.Printf("%T: %v\n", "Hello, 世界", "Hello, 世界")
	fmt.Printf("%T: %v\n", hello, hello)

	const myStringHello MyString = "Hello, 世界"

	fmt.Printf("%T: %v\n", myStringHello, myStringHello)
}
