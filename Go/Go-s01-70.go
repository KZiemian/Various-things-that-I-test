package main

import (
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("math: square root of negative number %g",
			f)
	}
	return math.Sqrt(f), nil
}

type NegativeSqrtError float64

func (f NegativeSqrtError) Error() string {
	return fmt.Sprintf("math: square root of negative number %g",
		float64(f))
}

func SqrtTwo(f float64) (float64, error) {
	if f < 0 {
		return 0, NegativeSqrtError(f)
	}
	return math.Sqrt(f), nil
}

func main() {
	x, err := SqrtTwo(1)
	fmt.Printf("x: %v, %#v, %T\n", x, x, x)
	fmt.Printf("err: %v,\n%#v,\n%T\n\n", err, err, err)

	x, err = SqrtTwo(0)
	fmt.Printf("x: %v, %#v, %T\n", x, x, x)
	fmt.Printf("err: %v,\n%#v,\n%T\n\n", err, err, err)

	x, err = SqrtTwo(-1)
	fmt.Printf("x: %v, %#v, %T\n", x, x, x)
	fmt.Printf("err: %v,\n%#v,\n%T\n", err, err, err)
}
