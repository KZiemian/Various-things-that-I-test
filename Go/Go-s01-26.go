package main

import "fmt"

func main() {
	var i int = 0
	var p *int = &i

	fmt.Printf("i: %v, %#v, %T\n", i, i, i)
	fmt.Printf("p: %v, %#v, %T\n", p, p, p)
}
