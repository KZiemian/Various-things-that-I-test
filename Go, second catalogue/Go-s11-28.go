package main

import "fmt"

type someType [3]interface {
	SomeMethod()
}

type someInt int

func (si someInt) SomeMethod() {
	fmt.Printf("someInt(%v)\n", si)
}

func someFunction[T someType](x T) {
	for _, v := range x {
		v.SomeMethod()
	}
}

func main() {
	var v someType

	v[0] = someInt(1)
	v[1] = someInt(2)
	v[2] = someInt(3)

	someFunction(v)
}
