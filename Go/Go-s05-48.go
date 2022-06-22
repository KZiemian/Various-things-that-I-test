package main

import (
	"fmt"
	"math"
)

type Calculator struct {
	acc float64
}

func (c *Calculator) Do(op func(float64) float64) float64 {
	c.acc = op(c.acc)

	return c.acc
}

func Add(n float64) func(float64) float64 {
	return func(acc float64) float64 {
		return acc + n
	}
}

func Sqrt() func(float64) float64 {
	return func(acc float64) float64 {
		return math.Sqrt(acc)
	}
}

func Cos() func(float64) float64 {
	return func(acc float64) float64 {
		return math.Cos(acc)
	}
}

func main() {
	var c Calculator

	fmt.Println(c.Do(Add(2)))
	fmt.Println(c.Do(Sqrt()))
	fmt.Println(c.Do(Cos()))
	fmt.Println(math.Cos(1.4142135623730951))
	fmt.Println(math.Sin(1.4142135623730951))
	fmt.Println(math.Cos(math.Sqrt2))
	fmt.Println(math.Sin(math.Sqrt2))
	fmt.Println(math.Pi/2)
}
