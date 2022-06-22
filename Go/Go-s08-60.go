package main

import (
	"os"
	"fmt"
	"log"
)

func main() {
	logger := log.New(os.Stderr, "Problem:", log.Lshortfile)

	logger.Fatal("Coś poszło nie tak.")
	fmt.Println("Działa")
}
