package main

import "fmt"

type myConstraint interface {
	scale(x float64)
}

type floatStruct struct {
	f float64
}

func (fStruct *floatStruct) scale(x float64) {
	fStruct.f = fStruct.f * x
}

func main() {
	f := &floatStruct{2}
	var varMyConstrains myConstraint

	fmt.Printf("f: %v, %T\n", f, f)
	fmt.Printf("varMyConstraint: %v, %T\n", varMyConstrains,
		varMyConstrains)

	varMyConstrains = f
	fmt.Printf("varMyConstraint: %v, %T\n", varMyConstrains,
		varMyConstrains)
}
