package main

import (
	"fmt"
	"time"
)

func main() {
	varTime := time.Now()

	fmt.Println("varTime:", varTime)

	fmt.Printf("\nvarTime: %v, %#v, %T\n", varTime, varTime,
		varTime)
	fmt.Printf("\nvarTime %%v: %v\n", varTime)
	fmt.Printf("\nvarTime %%#v: %#v\n", varTime)
	fmt.Printf("\nvarTime %%T: %T\n", varTime)
}
