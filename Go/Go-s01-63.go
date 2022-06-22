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
	fmt.Println("Println:", x)
	fmt.Printf("x: %v, %#v, %T\n", x, x, x)

	x = &Point{1, 2}
	fmt.Println("Println:", x)
	fmt.Printf("x: %v, %#v, %T\n", x, x, x)


	var y interface{}

	y = 2.4
	fmt.Println("Println:", y)
	fmt.Printf("y: %v, %#v, %T\n", y, y, y)

	y = &Point{1, 2}
	fmt.Println("Println:", y)
	fmt.Printf("y: %v, %#v, %T\n", y, y, y)
}
