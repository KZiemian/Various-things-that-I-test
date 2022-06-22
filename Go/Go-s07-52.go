package main

import (
	"fmt"
	"encoding/json"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	m := Message{"Alice", "Hello", 1294706395881547000}

	b, err := json.Marshal(m)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("b: %v\n", err)

	m1 := Message{}
	err = json.Unmarshal(b, &m1)
	fmt.Printf("m1: %v\n", m1)
	fmt.Printf("err: %v\n", err)
}
