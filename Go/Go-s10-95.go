package main

import "fmt"

func main() {
	fmt.Printf("funTest(int64(2)): %v\n", funTest(int64(2)))
}

func funTest[V int64|float64](x V) V {
	var r V
	// V := 2

	r = 2 * x
	// r += V

	return r
}
