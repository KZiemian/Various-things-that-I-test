package main

import "fmt"

func main() {
	a := []string{"John", "Paul"}
	b := []string{"George", "Ringo", "Pete"}

	fmt.Println("a:", a)
	fmt.Println("b:", b)

	a = append(a, b...)

	fmt.Println("a:", a)
}
