package main

import "fmt"

func main() {
	var u1 uint = 17
	var u2 = uint(17)
	u3 := uint(17)

	fmt.Printf("u1: %v, %T\n", u1, u1)
	fmt.Printf("u2: %v, %T\n", u2, u2)
	fmt.Printf("u3: %v, %T\n", u3, u3)
}
