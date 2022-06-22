package main

import "fmt"

func f() int {
	return 1
}

func main() {
	var y, z = 0, 2

	if x := f(); x < y {
		fmt.Println("x:", x)
	} else if x > z {
		fmt.Println("z:", z)
	} else {
		fmt.Println("y:", y)
	}
}
