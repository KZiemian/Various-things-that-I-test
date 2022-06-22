package main

import "fmt"

func main() {
	var x, y, res int = 1, 2, 0
	// var exprVar bool = true
	var exprVar bool = false

	fmt.Printf("x: %v; y: %v\n", x, y)

	if exprVar {
		res = x
	} else {
		res = y
	}

	fmt.Println("res:", res)
}
