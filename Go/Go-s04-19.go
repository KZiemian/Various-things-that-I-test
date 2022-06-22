package main

import "fmt"

func Insert(slice []int, index, value int) []int {
	slice = slice[0 : len(slice)+1]
	copy(slice[index+1:], slice[index:])
	slice[index] = value

	return slice
}

func main() {
	slice := make([]int, 10, 20)
	for i := range slice {
		slice[i] = i
	}
	fmt.Println(slice)
	slice = Insert(slice, 5, 99)
	fmt.Println(slice)
}
