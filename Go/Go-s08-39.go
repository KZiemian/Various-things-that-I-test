package main

import (
	"fmt"
	"log"
)

func main() {
	v := 3

	log.Fatalf("%s: %v\n", "Jakiś błąd", v)

	fmt.Println("Działa")
}
