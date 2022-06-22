package main

import (
	"fmt"
	"time"
)

// Publish prints text to stdout after the given time has expired.
// It dosen't block but returns right away.
func Publish(text string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println("BREAKING NEWS:", text)
	}()  // Note the parentheses. We must call the anonymous function.
}

func main() {
	Publish("A goroutine starts a new thread.", 5*time.Second)
	fmt.Println("Let's hope the news will published before I leave.")

	// Wait for the news to be published.
	time.Sleep(10 * time.Second)

	fmt.Println("Ten seconds later: I'm leaving now.")
}
