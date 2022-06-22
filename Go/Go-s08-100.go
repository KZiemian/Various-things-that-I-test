package main

import (
	"fmt"
	"os"
)

func main() {
	ids, err := os.Getgroups()
	if err != nil {
		fmt.Println("Coś poszło nie tak.")
		fmt.Println(err)
	}
	fmt.Println("ids:", ids)
}
