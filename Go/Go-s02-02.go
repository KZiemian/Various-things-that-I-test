package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	d1 := Student{"David", 1}
	d2 := Student{"David", 2}
	d3 := Student{"David", 1}

	fmt.Println("d1 == d2:", d1 == d2)
	fmt.Println("d1 == d3:", d1 == d3)
}
