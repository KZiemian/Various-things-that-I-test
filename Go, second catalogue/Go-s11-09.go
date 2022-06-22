package main

import (
	"fmt"
	"strings"
)

type Stringer interface {
	String() string
}

type StringableVector[T Stringer] []T

type someFloat float64

func (sF someFloat) String() string {
	return fmt.Sprintf("someFloat: %v", float64(sF))
}

func main() {
	var sliceSomeFloat StringableVector[someFloat] = make([]someFloat, 3)

	for i := 0; i < 3; i++ {
		sliceSomeFloat[i] = someFloat(i)
	}

	fmt.Println("sliceSomeFloat.String():", sliceSomeFloat.String())
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
