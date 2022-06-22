package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Printf("m: %v, %#v, %T\n\n", m, m, m)

	m["Answer"] = 42
	fmt.Printf("m[\"Answer\"]: %v, %#v, %T\n", m["Answer"],
		m["Answer"], m["Answer"])
	fmt.Printf("m: %v, %#v, %T\n\n", m, m, m)

	m["Answer"] = 48
	fmt.Printf("m[\"Answer\"]: %v, %#v, %T\n", m["Answer"],
		m["Answer"], m["Answer"])
	fmt.Printf("m: %v, %#v, %T\n\n", m, m, m)

	delete(m, "Answer")
	fmt.Printf("After delete, m[\"Answer\"]: %v, %#v, %T\n",
		m["Answer"], m["Answer"], m["Answer"])

	v, ok := m["Answer"]
	fmt.Printf("m[\"Answer\"]: %v, %#v, %T\n", v, v, v)
	fmt.Printf("ok: %v, %#v, %T\n", ok, ok, ok)
}
