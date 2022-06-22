package main

import "fmt"

type Foo struct {
	i int
	s string
}

type Bar int
type Qux []string

func main() {
	fooVar := Foo{i: 3, s: "Go"}
	barVar := Bar(0)
	quxVar := Qux([]string{"Martin", "Hairer", "to", "jest", "gość"})

	fmt.Printf("fooVar: %v, %T\n", fooVar, fooVar)
	fmt.Printf("barVar: %v, %T\n", barVar, barVar)
	fmt.Printf("quxVar: %v, %T\n", quxVar, quxVar)
}
