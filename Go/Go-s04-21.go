package main

import "fmt"

func Extend(slice []int, element int) []int {
	n := len(slice)
	if n == cap(slice) {
		newSlice := make([]int, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func Append(slice []int, items ...int) []int {
	for _, item := range items {
		slice = Extend(slice, item)
	}
	return slice
}

func main() {
	slice := []int{0, 1, 2, 3, 4}
	fmt.Println(slice)
	slice = Append(slice, 5, 6, 7, 8)
	fmt.Println(slice)
}
