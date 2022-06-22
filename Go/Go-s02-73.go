package main

import (
	"fmt"
	"time"
)

func Publish(text string, delay time.Duration) {
	go func() {
		start := time.Now()


		time.Sleep(delay)

		elapsedTime := time.Since(start)
		fmt.Println("elapsedTime:", elapsedTime)

		fmt.Println("BREAKING NEWS:", text)
	}()
}

func main() {
	start := time.Now()

	Publish("A goroutine starts a new thread.", 5*time.Second)
	fmt.Println("Let's hope the news will published before I leave.")


	time.Sleep(10 * time.Second)

	elapsedTime := time.Since(start)
	fmt.Println("elapsedTime:", elapsedTime)

	fmt.Println("Ten seconds later: I'm leaving now.")
}
