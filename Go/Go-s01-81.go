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
	var x interface{}

	fmt.Printf("x: %v, %#v, %T\n\n", x, x, x)

	x = 2.4
	fmt.Println(x)
	fmt.Printf("x: %v, %#v, %T\n\n", x, x, x)

	x = &Point{1, 2}
	fmt.Println(x)
	fmt.Printf("x: %v, %v, %T\n", x, x, x)
}
