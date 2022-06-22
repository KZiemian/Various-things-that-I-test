package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	fmt.Printf("a: %v, %#v, %T\n", a, a, a)
	fmt.Printf("a.Abs(): %v, %#v, %T\n\n", a.Abs(), a.Abs(), a.Abs())

	a = &v
	fmt.Printf("a: %v, %#v, %T\n", a, a, a)
	fmt.Printf("a.Abs(): %v, %#v, %T\n\n", a.Abs(), a.Abs(), a.Abs())

	fmt.Printf("v: %v, %#v, %T\n", v, v, v)
	fmt.Printf("v.Abs(): %v, %#v, %T\n", v.Abs(), v.Abs(), v.Abs())

	// a = v
	// Nie działa, bo Vertex nie implementuje metody Abs()
	// fmt.Println("a:", a)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
