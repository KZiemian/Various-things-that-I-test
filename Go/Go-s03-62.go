package main

import "fmt"

// func Sqrt(x float64) float64 {
// 	z := x / 2
// }

func main() {
	x := 2.0
	y := 2

	fmt.Println("x / 2:", x / 2)
	// fmt.Println("x/y:", x/y)
	fmt.Println("x / float64(y):", x / float64(y))
}
