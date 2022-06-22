package main

import (
	"fmt"
	"log"
)

func main() {
	v := 3
	log.Panicf("Jakiś błąd: %v\n", v)

	fmt.Println("Działa.")
}
