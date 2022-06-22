package main

import "fmt"

func f(a int) {
	// fmt.Println(a)

	result := make(chan error, 1)

	result <- nil

	return <-result
}


func main() {
	// f(0)
	fmt.Println("Here")
}
