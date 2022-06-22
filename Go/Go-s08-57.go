package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var (
		buf bytes.Buffer
		logger = log.New(&buf, "Jakiś problem:", log.Lshortfile)
	)

	logger.Fatal("Co za nonsens.")
	fmt.Println("Działa")
}
