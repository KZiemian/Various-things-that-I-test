package main

import (
	"os"
	"fmt"
	"log"
)

func main() {
	logger := log.New(os.Stderr, "Kolejny problem:", log.Lshortfile)

	fmt.Println("logger.Flags():", logger.Flags())
	fmt.Println("log.Flags():", log.Flags())
}
