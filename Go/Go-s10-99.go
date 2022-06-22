package main

import "fmt"

func main() {
	sliceVar1 := make([]int, 10)
	for i := 0; i < 10; i++ {
		sliceVar1[i] = 2 * i
	}
	// fmt.Println(sliceVar1)
	Print(sliceVar1)

	sliceVar2 := make([]float64, 10)
	for i := 0; i < 10; i++ {
		sliceVar2[i] = 2 / float64(i + 1)
	}

	Print(sliceVar2)

	sliceVar3 := []string{"Ian", "Lance", "Tylor"}
	Print(sliceVar3)
}


func Print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}
