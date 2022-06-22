package main

import "fmt"

func main() {
	fmt.Printf("f(2): %v\n", f(2))
}

func f(x float64) float64 {
	return 3
}
