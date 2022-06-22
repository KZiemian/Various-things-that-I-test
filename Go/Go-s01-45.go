package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()

	fmt.Printf("pos: %v, %#v, %T\n", pos, pos, pos)
	fmt.Printf("neg: %v, %#v, %T\n", neg, neg, neg)

	fmt.Println("pos", pos(0))
	fmt.Println(pos(1))
	fmt.Println("neg", neg(0))
	fmt.Println(neg(-1))

	fmt.Printf("pos: %v, %#v, %T\n", pos, pos, pos)
	fmt.Printf("neg: %v, %#v, %T\n", neg, neg, neg)
}
