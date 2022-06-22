package main

import "fmt"

func main() {
	var (
		i int
		b bool
		s string
	)
	n, err := fmt.Sscanln("33 true string", &i, &b, &s)
	fmt.Println("i:", i)
	fmt.Println("b:", b)
	fmt.Println("s:", s)
	fmt.Println("n:", n)
	if err != nil {
		fmt.Println("Coś poszło nie tak.")
	}
}
