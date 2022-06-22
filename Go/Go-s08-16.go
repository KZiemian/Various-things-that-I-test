package main

import "fmt"

func main() {
	data := struct {
		Title               string
		Firstname, Lastname string
		Rank                int
	}{
		"Commander", "Rob", "Pike", 10,
	}

	fmt.Printf("data: %v\n%T\n", data, data)
}
