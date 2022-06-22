package main

import "fmt"

const Pi = 3.14

func main() {
	fmt.Printf("Pi: %v, %#v, %T\n", Pi, Pi, Pi)
	Pi = 3.1415
	fmt.Printf("Pi: %v, %#v, %T\n", Pi, Pi, Pi)
}
