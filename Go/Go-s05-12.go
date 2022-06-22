package main

import "fmt"

type MyComplex128 complex128

const I = (0.0 + 1.0i)
const TypedI complex128 = (0.0 + 1.0i)

func main() {
	var mc MyComplex128
	mc = (0.0 + 1.0i)
	mc = I
	// mc = TypedI
	fmt.Println(mc)
}
