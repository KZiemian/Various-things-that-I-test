package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("function took: %v\n", time.Since(start))
}
