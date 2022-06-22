package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)
}
