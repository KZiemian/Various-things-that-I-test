package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	fmt.Printf("c: %v, %#v, %T\n", c, c, c)

	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Printf("i: %v, %#v, %T\n", i, i, i)
	}
}
