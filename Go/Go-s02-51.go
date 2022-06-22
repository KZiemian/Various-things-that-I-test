package main

import "fmt"

func a() int {
	i := 0
	defer fmt.Println("defer value of i:", i)
	i++
	return i
}

func main() {
	fmt.Println("a():", a())
}
