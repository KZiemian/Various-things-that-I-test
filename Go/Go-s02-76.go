package main

import (
	"fmt"
	"time"
)

func Publish(text string, delay time.Duration) (wait <-chan struct{}) {
	ch := make(chan struct{})
	go func() {
		time.Sleep(delay)
		fmt.Println(text)
		close(ch)
	}()

	return ch
}

func main() {
	wait := Publish("important news", 5*time.Second)
	// fmt.Println(wait)
	<-wait
}
