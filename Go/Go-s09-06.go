package main

import (
	"fmt"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Coś poszło nie tak.")
		fmt.Println(err)
	}
	fmt.Println("hostname:", hostname)
}
