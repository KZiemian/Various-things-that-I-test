package main

import "fmt"

func main() {
	fmt.Printf("%d\n", "hi")
	fmt.Printf("hi", "guys")
	fmt.Println()
	fmt.Printf("hi%d\n")
	fmt.Printf("%*s", 4.5, "hi\n")
	fmt.Printf("%.*s", 4.5, "hi\n")
	fmt.Printf("%*[2]d\n", 7)
	fmt.Printf("%.[2]d\n", 7)
}
