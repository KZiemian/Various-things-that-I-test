package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.8408},
}

func main() {
	fmt.Printf("m: %v, %#v, %T\n", m, m, m)
}
