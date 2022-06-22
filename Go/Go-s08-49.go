package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("log.Flags():", log.Flags())
	log.SetFlags(8)
	fmt.Println("log.Flags():", log.Flags())
}
