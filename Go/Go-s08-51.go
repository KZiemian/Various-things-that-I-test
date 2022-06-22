package main

import (
	"fmt"
	"log"
)

func main() {
	w := log.Writer()
	fmt.Printf("w: %v, %T\n", w, w)
}
