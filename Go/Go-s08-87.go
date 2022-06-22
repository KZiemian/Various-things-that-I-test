package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("Go-s08-86.go")
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
