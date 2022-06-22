package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stderr, "Jakiś błąd:", log.Lshortfile)

	logger.Println("Błąd")
	fmt.Println("Działa")
}
