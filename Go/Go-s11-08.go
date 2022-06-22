package main

import (
	"fmt"
	"strings"
)

type Stringer interface {
	String() string
}

type StringableVector[T fmt.Stringer] []T

type stupidFloat float64

// type stupidFloatOne struct {
// 	F float64
// }

// type stupidFloatTwo float64

func (sF stupidFloat) String() string {
	return fmt.Sprintf("Stupid float %v", float64(sF))
}

func main() {
	var varStupidFloat stupidFloat = -1.0

	sliceStupidFloat := make([]stupidFloat, 3)
	// sliceStuFloatTwo := make([]stupidFloatTwo, 5)

	for i := 0; i < 3; i++ {
		sliceStupidFloat[i] = stupidFloat(float64(i))
	}

	// for i := 0; i < 5; i++ {
	// 	sliceStuFloatTwo[i] = stupidFloatTwo(1)
	// }

	fmt.Println("stupidFloat.String(): ", varStupidFloat.String())

	fmt.Println("sliceStupidFloat:", sliceStupidFloat)
	// fmt.Println("sliceStuFloatTwo: ", sliceStuFloatTwo)

	fmt.Println("sliceStupidFloat:", sliceStupidFloat)
	fmt.Println("sliceStupidFloat.String():", sliceStupidFloat.String())
	// fmt.Println(sliceStupidFloat.String())
	// fmt.Println(sliceStuFloatTwo.String())
}

func (s StringableVector[T]) String() string {
	var sb strings.Builder

	for i, v := range s {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(v.String())
	}

	return sb.String()
}
