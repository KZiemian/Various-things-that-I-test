package main

import "fmt"

func main() {
	a := make([]int, 5)
	fmt.Printf("a: %v, %#v, %T\n", a, a, a)
	fmt.Printf("len(a) = %d, cap(a) = %d, %v\n\n",
		len(a), cap(a), a)

	b := make([]int, 0, 5)
	fmt.Printf("b: %v, %#v, %T\n", b, b, b)
	fmt.Printf("len(b) = %d, cap(b) = %d, %v\n\n",
		len(b), cap(b), b)

	c := b[:2]
	fmt.Printf("c: %v, %#v, %T\n", c, c, c)
	fmt.Printf("len(c) = %d, cap(c) = %d, %v\n\n",
		len(c), cap(c), c)

	d := c[2:5]
	fmt.Printf("d: %v, %#v, %T\n", d, d, d)
	fmt.Printf("len(d) = %d, cap(d) = %d, %v\n",
		len(d), cap(d), d)
}
