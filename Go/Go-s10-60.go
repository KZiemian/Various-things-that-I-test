package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	for i := 0; i < 4; i++ {
		fmt.Println("quote.Glass():", quote.Glass())
	}

	fmt.Println()

	for i := 0; i < 4; i++ {
		fmt.Println("quote.Go():", quote.Go())
	}

	fmt.Println()

	for i := 0; i < 4; i++ {
		fmt.Println("quote.Hello():", quote.Hello())
	}

	fmt.Println()

	for i := 0; i < 4; i++ {
		fmt.Println("quote.Opt():", quote.Opt())
	}
}
