package main

import (
	"fmt"
	"time"
)

func main() {
	// switch hour := time.Now().Hour(); {
	// case hour < 12:
	// 	fmt.Println("Good morning!")
	// case hour < 17:
	// 	fmt.Println("Good afternoon!")
	// default:
	// 	fmt.Println("Good evening!")
	// }

	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afernoon!")
	default:
		fmt.Println("Good evening!")
	}

	// fmt.Println("hour:", hour)
}
