package main

import "fmt"

func main() {
	v := 1

	for i := 0; i <= 20; i++ {
		fmt.Printf("2^%v = %v\n", i, v)
		v *= 2
	}
}
