package main

import (
	"fmt"
	"os"
)

func main() {
	exec, err := os.Executable()

	if err != nil {
		fmt.Println("Coś poszło nie tak.")
		fmt.Println("err:", err)
	}
	fmt.Printf("exec: %v\n", exec)
}
