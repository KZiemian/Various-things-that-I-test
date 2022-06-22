package main

import "fmt"

type MyString string

const typedHello string = "Hello, 世界"

func main() {
	var m MyString
	m = typedHello
	fmt.Println(m)
}
