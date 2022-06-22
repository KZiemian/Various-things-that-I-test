package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i <= 9; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(10 * time.Millisecond)
}
