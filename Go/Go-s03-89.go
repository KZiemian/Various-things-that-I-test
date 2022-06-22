package main

import "fmt"

func main() {
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	bOne := b[1:4]
	bTwo := b[:2]
	bThree := b[2:]
	bFour := b[:]

	fmt.Printf("b: %v, %T\n", b, b)
	fmt.Printf("bOne: %v, %T\n", bOne, bOne)
	fmt.Printf("bTwo: %v, %T\n", bTwo, bTwo)
	fmt.Printf("bThree: %v, %T\n", bThree, bThree)
	fmt.Printf("bFour: %v, %T\n", bFour, bFour)

	x := [3]string{"Лайка", "Белка", "Стрелка"}
	s := x[:]

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("s: %v, %T\n", s, s)
}
