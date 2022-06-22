package main

import "fmt"

func split(sum int) (x int, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	var x, y int = split(17)
	fmt.Printf("split(%v) = %v, %v\n", 17, x, y)
}
