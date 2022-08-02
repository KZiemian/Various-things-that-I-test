package main

import "fmt"

func main() {
	v := 'z'

	fmt.Printf("v: %T, %v\n", v, v)

	v = '.z'
	fmt.Printf("v: %T, %v\n", v, v)
}