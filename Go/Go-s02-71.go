package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("Hello from another goroutines")
	fmt.Println("Hello from main goroutine")

	time.Sleep(time.Second)
}
