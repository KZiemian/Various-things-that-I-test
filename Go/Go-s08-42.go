package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(log.Output(0, "Błąd 1"))
	fmt.Println(log.Output(1, "Błąd 2"))
	fmt.Println(log.Output(2, "Błąd 3"))
}
