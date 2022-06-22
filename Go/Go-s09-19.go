package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("Cos.txt")
	if err != nil {
		log.Println("Coś poszło nie tak.")
	}
	n, err := file.Write([]byte("Hello, Gophers!"))
	if err != nil {
		log.Println("Coś poszło nie tak.")
	}
	fmt.Println("n:", n)
}
