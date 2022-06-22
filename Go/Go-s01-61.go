package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7}

	for p := range primes {
		fmt.Println(p)
	}

	fmt.Println()

	for _, p := range primes {
		fmt.Println(p)
	}

	fmt.Println()

	for i, p := range primes {
		fmt.Printf("i: %v; p: %v;\n", i, p)
	}
}
