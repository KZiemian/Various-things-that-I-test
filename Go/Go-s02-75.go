package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "Hello!"
		close(ch)
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	v, ok := <-ch

	fmt.Printf("v: %v; ok: %v\n", v, ok)

	for v := range ch {
		fmt.Println(v)
	}
}
