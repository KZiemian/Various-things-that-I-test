package main

import "fmt"

type someInt[3] int

type someAnotherInt[T any] int

type someAnotherAnotherInt[_ any] int

type someYetAnotherInt[_] int

// type someOtherInt[float64 float64] int

// type someDifferentInt[float64] int

type someInterface[3] interface {
	Method()
}

// type someStruct[3 any, T any] struct {
// 	X int
// 	Y T
// }

func main() {
	var someIntVar someInt = [3]int{1, 2, 3}
	var arrayVar [2]int = [2]int{5, 6}

	fmt.Printf("someIntVar: %v, %T\n", someIntVar, someIntVar)
	fmt.Printf("arrayVar: %v, %T\n", arrayVar, arrayVar)
}
