package main

import "fmt"

type MySlice []int

func sumFun(s []int) int {
	var sum int = 0

	for _, v := range s {
		sum += v
	}

	return sum
}

func main() {
	var sliceMySlice MySlice = MySlice([]int{4, 5, 6})
	sliceInt := []int{1, 2, 3}

	fmt.Printf("sumFun(sliceInt) = %v\n", sumFun(sliceInt))
	fmt.Printf("sumFun(sliceMySlice) = %v\n", sumFun(sliceMySlice))
}
