package main

import (
	"fmt"
	"strconv"
)

type MyStringer interface {
	String() string
}

type Temp int

func (t Temp) String() string {
	return strconv.Itoa(int(t)) + " C"
}

type Point struct {
	x, y int
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func main() {
	var x MyStringer
	fmt.Printf("x: %v, %#v, %T\n", x, x, x)
	fmt.Println("x == nil:", x == nil)

	x = (*Point)(nil)
	fmt.Printf("x: %v, %#v, %T\n", x, x, x)
	fmt.Println("x == nil:", x == nil)
}
