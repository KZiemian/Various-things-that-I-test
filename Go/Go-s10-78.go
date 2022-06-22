package main

import "fmt"

func main() {
	factorial := 1
	fmt.Println(" 0! =", factorial)
	for i := 1; i <= 10; i++ {
		factorial *= i
		fmt.Printf("%2d! = %v\n", i, factorial)
	}
}
