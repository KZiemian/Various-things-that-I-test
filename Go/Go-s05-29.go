package main

import "fmt"

type Char byte

func main() {
	var c Char = '世'

	fmt.Printf("%v, %T\n", c, c)
}
