package main

import "fmt"

func Foo(n int) int {
	fmt.Println(n)
	return n
}

func main() {
	switch Foo(2) {
	case Foo(1), Foo(2), Foo(3):
		fmt.Println("First case")
	case Foo(4):
		fmt.Println("Second case")
	}
}
