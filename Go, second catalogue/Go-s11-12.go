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
	var varStringer Stringer = someFloat(7)
	sliceSomeFloat := make([]someFloat, 3)
	fmt.Printf("sliceSomeFloat type: %T\n", sliceSomeFloat)

	for i := 0; i < len(sliceSomeFloat); i++ {
		sliceSomeFloat[i] = someFloat(i)
	}

	var varStringableVector StringableVector[someFloat] = sliceSomeFloat

	fmt.Printf("someFloat type: %T\n", varStringer)
	fmt.Println("sliceSomeFloat:", sliceSomeFloat)
	fmt.Println("varStringableVector:", varStringableVector)
	fmt.Printf("varStringableVector type: %T\n", varStringableVector)
}
