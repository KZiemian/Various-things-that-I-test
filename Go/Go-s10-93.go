package main

import "fmt"

type myConstraint interface {
	int64
}

func main() {
	fmt.Printf("funTest(2): %v\n", funTest(2))
}

func funTest[V myConstraint](x V) V {
	var r V
	r = 2 * x

	return r
}
