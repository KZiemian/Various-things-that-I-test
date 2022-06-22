package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "Jakiś problemy:", log.Lshortfile)
	)

	logger.Print("Jakiś błąd")

	fmt.Print(&buf)

	fmt.Printf("logger: %v, %T\n\n\n", logger, logger)
	fmt.Printf("buf: %v, %T\n\n\n", buf, buf)
	fmt.Printf("&buf: %v, %T\n", &buf, &buf)
}
