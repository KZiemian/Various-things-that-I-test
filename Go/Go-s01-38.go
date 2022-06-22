package main

import "fmt"

func main() {
	pow := make([]int, 10)
	fmt.Printf("pow: %v, %#v, %T\n\n", pow, pow, pow)

	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	fmt.Printf("pow: %v, %#v, %T\n\n", pow, pow, pow)

	for i, value := range pow {
		fmt.Printf("pow[%d] = %v, %#v, %T\n",
			i, value, value, value)
	}
}
