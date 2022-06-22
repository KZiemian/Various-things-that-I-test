package main

import "fmt"

func main() {
	var m map[string]int
	m = make(map[string]int)

	m["route"] = 66
	i := m["route"]

	fmt.Println("m:", m)
	fmt.Println("i:", i)

	j := m["root"]
	fmt.Println("j:", j)

	n := len(m)
	fmt.Println("n:", n)

	delete(m, "route")

	fmt.Println("m:", m)

	i, ok := m["route"]
	fmt.Println("i:", i)
	fmt.Println("ok:", ok)

	_, ok = m["route"]
	fmt.Println("ok:", ok)
}
