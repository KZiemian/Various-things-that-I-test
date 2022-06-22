package main

import "fmt"

func main() {
	fmt.Printf("funTest(int64(2)): %v\n", funTest(int64(2)))
}

func funTest[V int64|float64](x V) V {
	var r V
	sliceTest := make([]V, 10)

	for i := 0; i < 10; i++ {
		sliceTest[i] = V(i)
	}

	for _, V := range sliceTest {
		r += V
	}

	var s V = x
	fmt.Printf("s: %v, %T\n", s, s)

	return r
}
