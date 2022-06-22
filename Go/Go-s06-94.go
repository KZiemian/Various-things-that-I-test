package main

import "fmt"

const (
	one1 = 1
	one2 = 1.000
	one3 = 1e3 - 99.0*10 - 9
	one4 = '\x01'
	one5 = '\u0001'
	one6 = 'b' - 'a'
	one7 = 1.0 + 3i - 3.0i
)

func main() {
	fmt.Printf("one1: %v, %T\n", one1, one1)
	fmt.Printf("one2: %v, %T\n", one2, one2)
	fmt.Printf("one3: %v, %T\n", one3, one3)
	fmt.Printf("one4: %v, %T\n", one4, one4)
	fmt.Printf("one5: %v, %T\n", one5, one5)
	fmt.Printf("one6: %v, %T\n", one6, one6)
	fmt.Printf("one7: %v, %T\n", one7, one7)
}
