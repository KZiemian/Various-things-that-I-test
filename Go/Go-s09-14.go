package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	dir, err := os.UserCacheDir()
	if err != nil {
		fmt.Println("Coś poszło nie tak.")
		log.Fatal(err)
	}
	fmt.Println("dir:", dir)
}
