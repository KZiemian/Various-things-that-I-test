package main

import "fmt"

func main() {
	fmt.Println("Test")
	fmt.Println("funTest(2):", funTest(2))
	fmt.Println("funTestSecond():", funTestSecond())
	fmt.Println("0xB3AD:", 0xB3AD)
}

func funTest(x int) int {
	return 3
}

func funTestSecond() int {
	return 33
}
