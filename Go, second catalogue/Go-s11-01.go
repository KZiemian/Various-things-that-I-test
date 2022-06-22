package main

import "fmt"

func main() {
	Print[int]([]int{1, 2, 3})
}

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}
