package main

import "fmt"

type T struct {
	n int
	f float64
	next *T
}

var (
	varInt int
	varFloat64 float64
	varBool bool
	varString string
	varIntPtr *int
	varT T
	varTArray [2]T
)

func main() {
	fmt.Println("varInt:", varInt)
	fmt.Printf("varInt: %v, %#v, %T\n\n", varInt, varInt, varInt)

	fmt.Println("varFloat64:", varFloat64)
	fmt.Printf("varFloat64: %v, %#v, %T\n\n", varFloat64, varFloat64,
		varFloat64)

	fmt.Println("varBool:", varBool)
	fmt.Printf("varBool: %v, %#v, %T\n\n", varBool, varBool, varBool)

	fmt.Println("varString:", varString)
	fmt.Printf("varString: %v, %#v, %T\n\n", varString, varString,
		varString)

	fmt.Println("varT:", varT)
	fmt.Printf("varT: %v, %#v, %T\n\n", varT, varT, varT)

	fmt.Println("[2]T{}:", [2]T{})
	fmt.Printf("[2]T{}: %v, %#v, %T\n\n", [2]T{}, [2]T{}, [2]T{})

	fmt.Println("varTArray:", varTArray)
	fmt.Printf("varTArray: %v, %#v, %T\n", varTArray, varTArray,
		varTArray)
}
