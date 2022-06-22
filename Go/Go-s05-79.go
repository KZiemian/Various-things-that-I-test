package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "?????"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

func main() {
	dog := Dog{}
	cat := Cat{}
	llama := Llama{}
	javaProgrammer := JavaProgrammer{}

	fmt.Println("Dog speak:", dog.Speak())
	fmt.Println("Cat speak:", cat.Speak())
	fmt.Println("Llama speak:", llama.Speak())
	fmt.Println("Java programmer speak:", javaProgrammer.Speak())
}
