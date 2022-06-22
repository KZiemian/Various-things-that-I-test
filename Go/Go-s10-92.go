package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	fmt.Printf("funTest(2): %v\n", funTest(2))
}

func funTest[V Number](x V) V {
	var r V
	r = 2 * x

	return r
}
