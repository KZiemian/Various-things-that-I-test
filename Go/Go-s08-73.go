package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stderr, "Jakiś błąd", log.Lshortfile)

	fmt.Println("logger.Flags():", logger.Flags())
	logger.SetFlags(3)
	fmt.Println("logger.Flags():", logger.Flags())
}
