package main

import (
	"fmt"
	// "time"
)

func main() {
	freeVar := "Hello "
	f := func(s string) {
		fmt.Println(freeVar + s)
	}
	f("Closuers")
	freeVar = "Goodbay "
	f("Clousers")
}
