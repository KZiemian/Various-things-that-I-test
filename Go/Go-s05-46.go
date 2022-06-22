package main

import (
	"fmt"
	"math"
)

type Calculator struct {
	acc float64
}

type opfunc func(float64, float64) float64

func (c *Calculator) Do(op opfunc, v float64) float64 {
	c.acc = op(c.acc, v)

	return c.acc
}

func Add(a, b float64) float64 {
	return a + b
}

func Sub(a, b float64) float64 {
	return a - b
}

func Mul(a, b float64) float64 {
	return a * b
}

func Sqrt(n, _ float64) float64 {
	return math.Sqrt(n)
}

func main() {
	var c Calculator

	fmt.Println(c.Do(Add, 5))
	fmt.Println(c.Do(Sub, 3))
	fmt.Println(c.Do(Mul, 8))
	fmt.Println(c.Do(Sqrt, 0))
}
