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

type Car struct {
	Vehicle
}

func main() {
	car := &Car{
		Vehicle{"Ford", 4},
	}

	fmt.Println(car.GetBrand())
	fmt.Println(car.GetSpeed())
}
