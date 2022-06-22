package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("math.Sqrt2:      %v, %#v, %T\n",
		math.Sqrt2, math.Sqrt2, math.Sqrt2)

	varSqrt2 := math.Sqrt2
	fmt.Printf("varSqrt2:        %v, %#v, %T\n", varSqrt2,
		varSqrt2, varSqrt2)

	var anotherVarSqrt2 float64 = math.Sqrt2
	fmt.Printf("anotherVarSqrt2: %v, %#v, %T\n",
		anotherVarSqrt2, anotherVarSqrt2, anotherVarSqrt2)
}
