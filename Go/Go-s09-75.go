package main

import "fmt"

func main() {
	for i := 1; i <= 25; i++ {
		fmt.Printf(" %%c, %d: %c\n", i, i)
		fmt.Printf(" %%q, %d: %q\n", i, i)
		fmt.Printf(" %%U, %d: %U\n", i, i)
		fmt.Printf("%%#U, %d: %#U\n", i, i)
	}
}
