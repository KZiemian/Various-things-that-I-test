package main

import "fmt"

type Student struct {
	Name string
}

func Bob(s Student) {
	s.Name = "Bob"
}

func Charlie(ps *Student) {
	ps.Name = "Charlie"
}

func main() {
	s := Student{"Alice"}

	Bob(s)
	fmt.Println("After Bob:", s)

	Charlie(&s)
	fmt.Println("After Charlie:", s)
}
