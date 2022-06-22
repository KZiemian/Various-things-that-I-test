package main

import "fmt"

func Extend(slice []int, element int) []int {
	n := len(slice)
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func main() {
	var iBuffer [10]int
	slice := iBuffer[0:0]
	for i := 0; i < 20; i++ {
		slice = Extend(slice, i)
		fmt.Println(slice)
	}
}
