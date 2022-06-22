package main

import (
	"fmt"
	"math"
	"errors"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	return math.Sqrt(f), nil
}

func main() {
	x, err := Sqrt(0)

	fmt.Printf("x: %v, %#v, %T\n", x, x, x)
	fmt.Printf("err: %v, %#v, %T\n\n", err, err, err)

	x, err = Sqrt(1)
	fmt.Printf("x: %v, %#v, %T\n", x, x, x)
	fmt.Printf("err: %v, %#v, %T\n\n", err, err, err)

	x, err = Sqrt(-1)
	fmt.Printf("x: %v, %#v, %T\n", x, x, x)
	fmt.Printf("err: %v, %#v, %T\n\n", err, err, err)

	if err != nil {
		fmt.Println(err)
	}
}
