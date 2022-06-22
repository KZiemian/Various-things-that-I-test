package main

import "fmt"

type Vehicle struct {
	brand  string
	wheels int
}

func (v *Vehicle) GetBrand() string {
	return v.brand
}

func (v *Vehicle) ComputeSpeed() int {
	return v.wheels * 5
}

type Car struct {
	Vehicle
}

type Motorcycle struct {
	Base Vehicle
}

type Speeder interface{
	GetSpeed() int
}

func printSpeed(s Speeder) {
	fmt.Println(s.GetSpeed())
}

func (c *Car) GetSpeed() int {
	return c.wheels * 3
}

func (m *Motorcycle) GetSpeed() int {
	return m.Base.ComputeSpeed()
}

func NewFord() *Car {
	return &Car{
		Vehicle{"Ford", 4},
	}
}

func main() {
	car := NewFord()


	printSpeed(car)
}
