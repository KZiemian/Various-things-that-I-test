package main

import "fmt"

const hello = "Hello, 世界"
const typedHello string = "Hello, 世界"

func main() {
	fmt.Printf("hello: %v, %T\n", hello, hello)
	fmt.Printf("typedHello: %v, %T\n", typedHello, typedHello)
}
