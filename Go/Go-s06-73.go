package main

import "fmt"

type MyString string

const hello = "Hello, 世界"

func main() {
	var m MyString

	m = "Hello, 世界"
	fmt.Println(m)
	fmt.Printf("%v, %T\n", m, m)

	m = hello
	fmt.Println(m)
	fmt.Printf("%v, %T\n", m, m)
}
