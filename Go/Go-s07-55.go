package main

import "fmt"

func ack(m, n int) int {
	if m == 0 {
		return n + 1
	} else if n == 0 {
		return ack(m - 1, 1)
	} else {
		return ack(m - 1, ack(m, n - 1))
	}
}

func main() {
	fmt.Printf("ack(4, 1) is: %v\n", ack(4, 1))
}
