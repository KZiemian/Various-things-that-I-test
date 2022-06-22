package main

import "fmt"

func Map(f func(v int) bool, vs []int) []bool {
	if len(vs) == 0 {
		return nil
	}

	return append(
		Map(f, vs[:len(vs)-1]),
		f(vs[len(vs)-1]))
}

func main() {
	isEven := func(v int) bool {
		return v%2 == 0
	}

	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(Map(isEven, nums))
}
