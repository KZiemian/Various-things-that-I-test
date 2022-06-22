package main

import (
	"os"
	"fmt"
	"log"
)

func main() {
	logger := log.New(os.Stderr, "Jakiś błąd:", log.Lshortfile)

	logger.Panicln("Błąd")
	fmt.Println("Działa")
}
