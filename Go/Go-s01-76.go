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
	v := Vertex{3, 4}

	a = v
	fmt.Printf("a: %v, %#v, %T\n", a, a, a)
	fmt.Printf("a.Abs(): %v, %#v, %T\n", a.Abs(), a.Abs(), a.Abs())

	a = &v
	fmt.Println("Działa, bo do zbioru metod *T należą wszystkie metody",
		"zdefioniowane dla typu T.")
	fmt.Printf("a: %v, %#v, %T\n", a, a, a)
	fmt.Printf("a.Abs(): %v, %#v, %T\n", a.Abs(), a.Abs(), a.Abs())

	// fmt.Println(a)
	// fmt.Println(math.Sqrt2)
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
