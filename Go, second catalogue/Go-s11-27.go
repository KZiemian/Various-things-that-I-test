package main

import "fmt"

func funcOne(x int, x int) int {
	return 2 * x
}

// func funcTwo(x int, x float64) float64 {
// 	return 2 * x
// }

// func funcThree(x int, x float64) int {
// 	return 2 * x
// }

func main() {
	fmt.Println("funcOne(2, 2):", funcOne(2, 2))
	// fmt.Println("funcTwo(3, 3):", funcTwo(3, 3))
	// fmt.Println("funcThree(4, 4):", funcThree(4, 4))
}
