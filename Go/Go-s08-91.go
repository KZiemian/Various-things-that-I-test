package main

import (
	"fmt"
	"os"
)

func main() {
	environment := os.Environ()

	for _, v := range environment {
		fmt.Println(v)
	}
}
