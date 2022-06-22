package main

import (
	"os"
	"fmt"
	"log"
)

func main() {
	logger := log.New(os.Stderr, "Jakiś problem", log.Lshortfile)
	v := 3

	logger.Fatalf("Zmienna: %v\n", v)
	fmt.Println()
}
