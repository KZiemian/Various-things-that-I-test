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

// ack(1, 0) = ack(0, 1) = 2
// ack(0, 1) = 2
// ack(1, 1) = ack(0, ack(1, 0)) = 3
// ack(2, 0) = ack(1, 1) = 3
// ack(1, 2) = ack(0, ack(1, 1)) = 4
// ack(2, 1) = ack(1, ack(2, 0)) = ack(1, 3) = ack(0, ack(1, 2)) = 5

func main() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("ackerman(%v, %v) is: %v\n", i, j,
				ack(i, j))
		}
	}
}
