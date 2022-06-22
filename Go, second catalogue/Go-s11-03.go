package main

import "fmt"

func main() {
	sliceVar1 := make([]int, 10)
	for i := 0; i < 10; i++ {
		sliceVar1[i] = 2 * i
	}
	Print[int](sliceVar1)
}

func Print[T interface{}](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}
