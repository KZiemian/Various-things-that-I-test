package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "Głupoty:", log.Lshortfile)
	)

	logger.Print("Jakaś informacja")

	fmt.Print(&buf)
}
