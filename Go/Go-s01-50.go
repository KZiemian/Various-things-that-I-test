package main

import "fmt"

func main() {
	fmt.Printf("nil: %v, %#v, %T\n", nil, nil, nil)

	var a interface{} = nil
	fmt.Printf("a: %v, %#v, %T\n", a, a, a)
}
