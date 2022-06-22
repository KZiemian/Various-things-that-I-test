package main

import (
	"fmt"
	"time"
)

const (
	timeout = 500 * time.Millisecond
	sqrt2 = 1.414213
)

func main() {
	fmt.Println("timeout:", timeout)
	fmt.Println("sqrt2:", sqrt2)

	fmt.Printf("timeout: %v, %T\n", timeout, timeout)
	fmt.Printf("sqrt2: %v, %T\n", sqrt2, sqrt2)


	// Tu jest błąd
	// var x float64 = timeout

	// fmt.Println("x:", x)
	// fmt.Printf("x: %v, %T\n", x, x)
}
