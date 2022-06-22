package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic says:", r)
		}
	}()

	panic("Panicking!")

	fmt.Println("Here")
}
