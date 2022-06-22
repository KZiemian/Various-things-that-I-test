package main

import "fmt"

type MyString string

const hello = "Hello, 世界"
const myStringHello MyString = "Hello, 世界"

func main() {
	fmt.Printf("%T: %v\n", "Hello, 世界", "Hello, 世界")
	fmt.Printf("%T: %v\n", hello, hello)
	fmt.Printf("%T: %v\n", myStringHello, myStringHello)
}
