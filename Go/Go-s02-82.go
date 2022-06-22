package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	fmt.Printf("t: %v, %#v, %T\n", t, t, t)
}
