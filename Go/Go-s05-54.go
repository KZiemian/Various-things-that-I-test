package main

import "fmt"

func main() {
	n := 4
	n2 := square(n)
	fmt.Println(n2)
}

func square(x int) int {
	return x * x
}
