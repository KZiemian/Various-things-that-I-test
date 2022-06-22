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
	fmt.Printf("%v, %#v, %T\n", x, x, x)

	x = Temp(24)
	fmt.Printf("%v, %#v, %T\n", x, x, x)

	x = &Point{1, 2}
	fmt.Printf("%v, %#v, %T\n", x, x, x)

	x = (*Point)(nil)
	fmt.Printf("%v, %#v, %T\n", x, x, x)
}
