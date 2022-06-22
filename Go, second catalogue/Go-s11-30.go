package main

import "fmt"

type MySlice []int

type MyInteger interface {
	int | int64 | int32 | int16 | int8
}

func Double[E MyInteger](s []E) []E {
	fmt.Printf("s value: %v\n", s)
	fmt.Printf("s value: %T\n\n", s)

	r := make([]E, len(s))

	for i, v := range s {
		r[i] = v + v
	}

	return r
}

func main() {
	var sliceMySlice MySlice = MySlice([]int{1, 2, 3})
	var V1 = Double(sliceMySlice)

	fmt.Printf("sliceMySlice value: %v\n", sliceMySlice)
	fmt.Printf("sliceMySlice type:  %T\n\n", sliceMySlice)

	fmt.Printf("V1 value: %v\n", V1)
	fmt.Printf("V1 type:  %T\n", V1)
}
