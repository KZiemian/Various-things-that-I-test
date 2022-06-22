package main

import (
	"fmt"
	"time"
	"math"
)

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(math.Sqrt2)
	}
	fmt.Printf("function took %v\n", time.Since(start))
}
