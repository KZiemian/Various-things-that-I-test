package main

import "fmt"

func funTest[V int64 | float64](x V) V {
	var s V
	sliceTest := make([]V, 10)

	for i := 0; i < 10; i++ {
		sliceTest[i] = V(i)
	}

	for _, V := range sliceTest {
		s += V
		fmt.Printf("V type: %T\n", V)
	}

	return s
}

func main() {
	fmt.Printf("funTest(2): %v\n", funTest(int64(2)))
}
