package main

import "fmt"

const typedHello string = "Hello, 世界"

type MyString string

func main() {
	var m MyString
	m = MyString(typedHello)
	fmt.Println(m)
}
