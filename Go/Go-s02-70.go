package main

import "fmt"

func main() {
	go fmt.Println("Hello from another goroutine")
	fmt.Println("Hello from main goroutine")
}
