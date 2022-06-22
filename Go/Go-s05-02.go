package main

import "fmt"

const hello = "Hello, 世界"

type MyString string

func main() {
	var m MyString
	m = hello
	fmt.Println(m)

	var n MyString
	n = "Hello, 世界"
	fmt.Println(n)
}
