package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}


	var z_prev  float64 = x
	var z_now   float64 = x
	var delta_z float64 = 1.0

	for i := 0; i < 20; i++ {
		z_prev = z_now

		z_now -= (z_now*z_now - x) / (2*x)

		delta_z = math.Abs(z_now - z_prev)

		if delta_z < 0.000001 {
			break
		}
	}

	return z_now, nil
}

func main() {
	fmt.Println("math.Sqrt2:", math.Sqrt2)
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
