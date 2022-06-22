package main

import "fmt"

func main() {
	fmt.Printf("0i: %v, %T\n", 0i, 0i)
	fmt.Printf("0123i: %v, %T\n", 0123i, 0123i)
	fmt.Printf("0o123i: %v, %T\n", 0o123i, 0o123i)
	fmt.Printf("0xabci: %v, %T\n", 0xabci, 0xabci)
	fmt.Printf("0.i: %v, %T\n", 0.i, 0.i)
	fmt.Printf("2.71828i: %v, %T\n", 2.71828i, 2.71828i)
	fmt.Printf("1.e+0i: %v, %T\n", 1.e+0i, 1.e+0i)
	fmt.Printf("6.67428e-11i: %v, %T\n", 6.67428e-11i, 6.67428e-11i)
	fmt.Printf("1E6i: %v, %T\n", 1E6i, 1E6i)
	fmt.Printf(".25i: %v, %T\n", .25i, .25i)
	fmt.Printf(".12345E5i: %v, %T\n", .12345E5i, .12345E5i)
	fmt.Printf("0x1p-2i: %v, %T\n", 0x1p-2i, 0x1p-2i)
}
