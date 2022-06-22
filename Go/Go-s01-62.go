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

	x = Temp(24)
	fmt.Println(x.String()) // 24 ^\circ C

	x = &Point{1, 2}
	fmt.Println(x.String()) // (1,2)

	x = Temp(24)
	fmt.Println(x) // 24 ^\cirs C

	x = &Point{1, 2} // (1,2)
	fmt.Println(x)


	var y interface{}

	fmt.Printf("\n\ny: %v, %#v, %T\n", y, y, y)

	y = 2.4
	fmt.Println(y)

	y = &Point{1, 2}
	fmt.Println(y)
}
