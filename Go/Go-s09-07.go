package main

import (
	"fmt"
	"os"
)

func main() {
	v := os.IsPathSeparator('/')
	fmt.Println("os.IsPathSeparator('/'):", v)
	v = os.IsPathSeparator('\\')
	fmt.Println("os.IsPathSeparator('\\'):", v)
}
