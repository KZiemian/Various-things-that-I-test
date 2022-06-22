package main

import "fmt"

func main() {
	const MaxUint = ^uint(0)
	fmt.Printf("%x\n", MaxUint)
	fmt.Printf("%v, %T\n", MaxUint, MaxUint)
}
