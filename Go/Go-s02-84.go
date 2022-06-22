package main

import (
	"fmt"
	// "time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// time.Sleep(2*time.Second)
	close(c)
}

func main() {
	c := make(chan int, 10)

	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
