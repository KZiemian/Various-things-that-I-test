package main

import "fmt"

type someNumber interface {
	int | float64
}

type someOtherNumber interface {
	int | int64
}

// func someFunction[T someNumber, T someOtherNumber](x T) T {
// 	return 2 * x
// }

// func someFunction[T someNumber](x T) T {
// 	return 2 * x
// }

// func someFunction[T someOtherNumber](x T) T {
// 	return 2 * x
// }

func someFunction[T someNumber, T someNumber](x T) T {
	return 2 * x
}

func main() {
	fmt.Println("someFunction(2):", someFunction(2))
}
