package main

import "fmt"

func main() {
	var u uint = 0
	var i int = 1

	fmt.Println("i + int(u):", i + int(u))
	fmt.Println("uint(i) + u:", uint(i) + u)
}
