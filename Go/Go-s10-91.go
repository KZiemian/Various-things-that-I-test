package main

import "fmt"

func main() {
	fmt.Printf("funTest(2): %v\n", funTest(2))
}

func funTest[V int64](x V) V {
	var r V
	r = 2*x

	return r
}
