package main

import "fmt"

func main() {
	str1 := "Hello, 世界"

	fmt.Printf("%v, %T\n", str1, str1)

	var str2 = "Hello, 世界"

	fmt.Printf("%v, %T\n", str2, str2)

	var str3 string = "Hello, 世界"

	fmt.Printf("%v, %T\n", str3, str3)
}
