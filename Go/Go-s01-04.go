package main

import (
	"fmt"
	"math/rand"
)


func main() {
	fmt.Println("rand.Seed:", rand.Seed)
	fmt.Printf("rand.Seed: %v, %#v, %T\n", rand.Seed, rand.Seed,
		rand.Seed)
}
