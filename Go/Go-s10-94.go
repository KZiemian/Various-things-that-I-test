package main

import "fmt"

func main() {
	var varInt int64 = 2
	var varFloat float64 = 2.0

	fmt.Printf("funTest(varInt): %v\n", funTest(varInt))
	fmt.Printf("funTest(varFloat): %v\n", funTest(varFloat))
}

func funTest[V int64|float64](x V) V {
	var r V
	r = 2 * x

	// fmt.Printf("V type: %v, %T\n", V, V)
	fmt.Printf("r: %v, %T\n", r, r)

	return r
}
