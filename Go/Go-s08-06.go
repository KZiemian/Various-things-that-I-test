package main

import "fmt"

func main() {
	n := 4
	inc(&n)

	fmt.Println(n)
}

func inc(x *int) {
	*x++
}
