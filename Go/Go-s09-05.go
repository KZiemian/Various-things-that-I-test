package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Coś poszło nie tak.")
		fmt.Println(err)
	}
	fmt.Println("dir:", dir)
}
