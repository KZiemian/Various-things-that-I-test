package main

import "fmt"

func main() {
	sliceOne   := make([]int, 0)
	sliceTwo   := make([]float64, 0)
	sliceThree := make([]int, 1)
	sliceFour  := make([]float64, 1)

	fmt.Printf("sliceOne: %v, %v, %v\n", sliceOne, len(sliceOne),
		cap(sliceTwo))
	fmt.Printf("sliceTwo: %v, %v, %v\n", sliceTwo, len(sliceTwo),
		cap(sliceTwo))
	fmt.Printf("sliceThree: %v, %v, %v\n", sliceThree, len(sliceThree),
		cap(sliceThree))
	fmt.Printf("sliceFour: %v, %v, %v\n", sliceFour, len(sliceFour),
		cap(sliceFour))

	sliceOne = append(sliceOne, 1, 2, 3)
	fmt.Println("After append.")
	fmt.Printf("sliceOne: %v, %v, %v\n", sliceOne, len(sliceOne),
		cap(sliceOne))
}
