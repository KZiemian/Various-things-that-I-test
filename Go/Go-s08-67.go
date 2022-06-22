package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stderr, "Jakiś błąd", log.Lshortfile)
	v := 3

	logger.Panicf("Zmienna: %v\n", v)

	fmt.Println("Działa")
}
