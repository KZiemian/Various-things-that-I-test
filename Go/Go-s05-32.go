package main

import "fmt"

const InfiniteOnes = ^0

func main() {
	fmt.Println(InfiniteOnes)
	fmt.Printf("%v, %T\n", InfiniteOnes, InfiniteOnes)
}
