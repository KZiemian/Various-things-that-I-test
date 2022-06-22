package main

import (
	"os"
	"fmt"
	"log"
)

func main() {
	logger := log.New(os.Stderr, "Jakiś problem:", log.Lshortfile)

	logger.Fatalln("Coś poszło nie tak.")
	fmt.Println("Działa.")
}
