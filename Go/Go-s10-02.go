package main

import "fmt"

func main() {
	s := "string"
	for _, v := range s {
		fmt.Printf("%v, %T\n", v, v)
	}
}
