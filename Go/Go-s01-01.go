package main

import "fmt"

func main() {
	var string1 string = "Hello, 世界"

	fmt.Println(string1)

	fmt.Printf("\nstring1: %v, %#v, %T\n", string1, string1,
		string1)
	fmt.Printf("\nstring1 %%v: %v\n", string1)
	fmt.Printf("\nstring1 %%#v: %#v\n", string1)
	fmt.Printf("\nstring1 %%T: %T\n", string1)
}
