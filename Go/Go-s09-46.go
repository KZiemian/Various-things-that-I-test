package main

import "fmt"

func main() {
	var i1 uint = 1000000

	fmt.Printf("%%v: %v\n", i1)
	fmt.Printf("%%#v: %#v\n", i1)
	fmt.Printf("%%+v: %+v\n", i1)
	fmt.Printf("%%d: %d\n", i1)
	fmt.Printf("%%x: %x\n", i1)
	fmt.Printf("%%#x: %#x\n", i1)

	f1 := 1234.56789
	fmt.Printf("%%f: %f\n", f1)
	fmt.Printf("%%9f: %9f\n", f1)
	fmt.Printf("%%.2f: %.2f\n", f1)
	fmt.Printf("%%9.2f: %9.2f\n", f1)
	fmt.Printf("%%9.f: %9.f\n", f1)

	var i interface{} = 23
	fmt.Printf("%v\n", i)
}
