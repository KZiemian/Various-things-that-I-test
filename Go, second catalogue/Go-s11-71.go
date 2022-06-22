package main

import "fmt"

func main() {
	v := 1.0
	gain := 1.1

	for i := 0; i < 10; i++ {
		v *= gain
	}

	fmt.Printf("v = %v\n", v)
}
