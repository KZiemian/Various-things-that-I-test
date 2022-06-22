package main

import "fmt"

type MyFloat64 float64
const Zero = 0.0
const TypedZeor float64 = 0.0

func main() {
	var f32 float32
	f32 = 0.0
	f32 = Zero
	// f32 = TypedZero
	fmt.Println(f32)
}
