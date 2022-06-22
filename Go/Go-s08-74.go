package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stderr, "Jakiś błąd:", log.Lshortfile)

	logger.Println("Błąd 1")
	logger.SetPrefix("Inny błąd:")
	logger.Println("Błąd 2")

	fmt.Println("Działa")
}
