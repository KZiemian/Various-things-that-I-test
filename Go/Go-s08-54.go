package main

import (
	"fmt"
	"log"
)

func main() {
	v := log.Default()

	fmt.Printf("%v, %T\n", v, v)
}
