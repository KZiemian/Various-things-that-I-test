package main

import "fmt"

func main() {
	v := 1.0
	gain := 1.5

	for i := 0; i < 10; i++ {
		fmt.Printf("v = %v\n", v)
		v *= gain
	}
	fmt.Printf("v = %v\n", v)
}
