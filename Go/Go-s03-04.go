package main

import "fmt"

func main() {
	channelInt := make(chan int, 1)

	channelInt <- 0
	fmt.Println("<-channelInt:", <-channelInt)
	// fmt.Println(channelInt)
}
