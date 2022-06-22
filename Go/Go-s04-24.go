package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	slice2 := []int{55, 66, 77}
	fmt.Println("Start slice:", slice)
	fmt.Println("Start slice2:", slice2)

	slice = append(slice, 4)
	fmt.Println("Add one item:", slice)

	slice = append(slice, slice2...)
	fmt.Println("Add one slice:", slice)

	slice3 := append([]int(nil), slice...)
	fmt.Println("Copy a slice:", slice3)

	fmt.Println("Before append to self:", slice)
	slice = append(slice, slice...)
	fmt.Println("After append to self:", slice)
}
