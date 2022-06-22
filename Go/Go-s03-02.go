package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	fmt.Printf("%v\n", timer1)
	fmt.Printf("%#v\n", timer1)
	fmt.Printf("%T\n", timer1)
}
