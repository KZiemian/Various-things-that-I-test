package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i <= 20; i++ {
		rand.Seed(int64(i))
		fmt.Printf("rand.Seed(%v), rand.Intn(10) = %v\n", i,
			rand.Intn(10))
	}
}
