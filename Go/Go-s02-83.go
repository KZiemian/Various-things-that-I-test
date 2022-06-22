package main

import (
	"fmt"
	"time"
)

func main() {
	var t time.Time = time.Now()

	fmt.Printf("t: %v,\n", t)
	fmt.Printf("   %#v,\n", t)
	fmt.Printf("   %T\n", t)
}
