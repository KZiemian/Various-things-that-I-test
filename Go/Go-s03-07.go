package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	fmt.Println("t, 1:", t)
	t = time.Now()
	fmt.Println("t, 2:", t)
	time.Sleep(2 * time.Second)
	t = time.Now()
	fmt.Println("t, 3:", t)
}
