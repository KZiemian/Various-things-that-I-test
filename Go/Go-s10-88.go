package main

import "fmt"

func funTest[V float64 | int64](x V) V {
	var v V
	v = 2 * x

	return v
}

func main() {
	fmt.Printf("funTest(2): %v\n", funTest(2))
	// ./Go-s10-88.go:6:40: int does not implement float64|int64
}
