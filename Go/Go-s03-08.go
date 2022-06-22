package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	time.Sleep(2 * time.Second)
	elapsedTime := time.Since(t)
	fmt.Printf("elapsedTime: %v, %#v, %T\n", elapsedTime,
		elapsedTime, elapsedTime)
}
