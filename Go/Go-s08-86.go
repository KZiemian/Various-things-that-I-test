package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("file.go")
	if err != nil {
		log.Fatal(err)
	}
}
