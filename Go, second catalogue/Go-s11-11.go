package main

import "fmt"

type someFloat float64

func main() {
	var floatVar float64 = 1.0
	var someFloatVar someFloat = floatVar

	fmt.Println(someFloatVar)
}
