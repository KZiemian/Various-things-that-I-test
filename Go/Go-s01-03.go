package main

import (
	"fmt"
	"math/rand"
)

func main() {
	varRandIntn := rand.Intn(10)

	fmt.Println("varRandIntn:", varRandIntn)
	fmt.Printf("varRandIntn: %v, %#v, %T\n", varRandIntn, varRandIntn,
		varRandIntn)
	fmt.Printf("\nvarRandIntn %%v: %v\n", varRandIntn)
	fmt.Printf("\nvarRandIntn %%#v: %#v\n", varRandIntn)
	fmt.Printf("\nvarRandIntn %%T: %T\n", varRandIntn)
}
