package main

import "fmt"

func main() {
	var (
		a int
		b int
		c int
	)
	n, err := fmt.Scan(&a, &b, &c)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("n:", n)
	if err != nil {
		fmt.Println("Coś poszło nie tak.")
	}
}
