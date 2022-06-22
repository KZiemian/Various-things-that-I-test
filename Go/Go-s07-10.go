package main

import "fmt"

type Vehicle struct {
	brand  string
	wheels int
}

func (v *Vehicle) GetBrand() string {
	return v.brand
}

func (v *Vehicle) GetSpeed() int {
	return v.wheels * 5
}

type Motorcycle struct {
	Base Vehicle
}

func main() {
	motorcycle := &Motorcycle{
		Vehicle{"BMV", 2},
	}

	fmt.Println(motorcycle.Base.GetBrand())
	fmt.Println(motorcycle.Base.GetSpeed())
}
