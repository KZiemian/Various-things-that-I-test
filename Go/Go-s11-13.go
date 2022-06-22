package main

import "fmt"

type Stringer interface {
	String() string
}

type StringableVector[T Stringer] []T

type someFloat float64

func (sF someFloat) String() string {
	return fmt.Sprintf("someFloat: %v", float64(sF))
}

func main() {
	var stringerVar Stringer = someFloat(7)
	sliceSomeFloat := make([]someFloat, 3)
	var stringableVectorVar StringableVector[someFloat] = sliceSomeFloat

	fmt.Printf("stringerVar type: %T\n", stringerVar)
	fmt.Printf("sliceScomeFloat type: %T\n", sliceSomeFloat)
	fmt.Printf("stringableVectorVar type: %T\n", stringableVectorVar)
}
