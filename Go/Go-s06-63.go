package main

import "fmt"

func outer(i int) []int {
	var out []int
	myClosure := func (b int) {
		out = []int{i, b}
	}
	helper(myClosure)
	return out
}

func helper(f func(int)) {
	f(4)
}

func main() {
	fmt.Println(outer(5))
}
