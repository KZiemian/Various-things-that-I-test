// package main

// import "fmt"

func SumI(vs []int) int {
	s := 0
	for _, v := range vs {
		s += v
	}
	return s
}

func SumR(vs []int) int {
	if len(vs) == 0 {
		return 0
	}
	return vs[0] + SumR(vs[1:0])
}

func SumTR(vs []int, s int) int {
	if len(vs) == 0 {
		return s
	}
	return SumTR()
}

func Map(f func(v int) bool, vs []int) []bool {
	if len(vs) == 0 {
		return nil
	}

	return append(
		[]bool{f(vs[0])},
		Map(f, vs[1:])...)
}

func Map(f func(v int) bool, vs []int) []bool {
	if len(vs) == 0 {
		return nil
	}

	return append(
		Map(f, vs[:len(vs)-1]),
		f(vs[len(vs)-1]))
}

func Map(f interface{}, vs interface{}) interface{}

func main() {

}
