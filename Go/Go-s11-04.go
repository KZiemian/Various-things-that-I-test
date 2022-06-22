package main

import "fmt"

type Stringer interface {
	String() string
}

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("Współrzędne punktu: %v, %v", p.X, p.Y)
}

func main() {
	slicePoints := make([]Point, 10)
	for i := 0; i < 10; i++ {
		slicePoints[i] = Point{i, i + 3}
	}

	for i, v := range slicePoints {
		fmt.Printf("slicePoints[%v]: %v\n", i, v)
	}

	sliceString := Stringify(slicePoints)
	fmt.Printf("sliceString: %T\n", sliceString)
	fmt.Printf("len(sliceString): %v\n", len(sliceString))
	fmt.Printf("cap(sliceString): %v\n", cap(sliceString))

	for _, v := range sliceString {
		fmt.Println(v)
	}
}

func Stringify[T Stringer](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}

	return ret
}
