package main

import "fmt"

func main() {
	var i int = 42
	var f float64 = float64(i)
	var u uint = f

	fmt.Printf("i: %v, %#v, %T\n", i, i, i)
	fmt.Printf("f: %v, %#v, %T\n", f, f, f)
	fmt.Printf("u: %v, %#v, %T\n", u, u, u)
}
