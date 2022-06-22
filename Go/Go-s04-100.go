package main

import "fmt"

type MyString string

const myStringHello MyString = "Hello, 世界"

func main() {
	var m MyString
	m = myStringHello
	fmt.Println(m)
}
