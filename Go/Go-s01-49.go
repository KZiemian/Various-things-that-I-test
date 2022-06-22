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

	fmt.Printf("a: %v, %#v, %T\n", a, a, a)
	fmt.Printf("f: %v, %#v, %T\n", f, f, f)

	a = f
	fmt.Println("Po a = f")
	fmt.Printf("a: %v, %#v, %T\n", a, a, a)
	a = &v
	fmt.Println("Po a = &v")
	fmt.Printf("a: %v, %#v, %T\n", a, a, a)

	// a = v
	// To jest błąd, bo typ Vertex nie implementuje
	// metody Abs
	fmt.Println("a.Abs():", a.Abs())
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
