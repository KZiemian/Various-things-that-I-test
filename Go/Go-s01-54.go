package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	fmt.Printf("ch: %v, %#v, %T\n", ch, ch, ch)

	ch <- 1
	ch <- 2
	fmt.Printf("ch: %v, %#v, %T\n", ch, ch, ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
