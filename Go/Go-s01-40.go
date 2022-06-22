package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	fmt.Printf("m: %v, %#v, %T\n", m, m, m)

	m = make(map[string]Vertex)
	fmt.Printf("m: %v, %#v, %T\n", m, m, m)

	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Printf("m: %v, %#v, %T\n", m, m, m)
	fmt.Println("m[\"Bell Labs\"]:", m["Bell Labs"])
}
