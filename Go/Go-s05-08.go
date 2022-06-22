package main

import "fmt"

type MyFloat64 float64

const Zero = 0.0
const TypedZero float64 = 0.0

func main() {
	var mf MyFloat64
	mf = 0.0
	mf = Zero
	// mf = TypedZero
	fmt.Println(mf)
}
