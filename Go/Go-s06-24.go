//usr/local/go/bin/go run "$0" "$@"

package main

import (
	"fmt"
	"os"
)

func main() {
	s := "world"

	if len(os.Args) > 1 {
		s = os.Args[1]
	}
	fmt.Printf("Hello, %v!\n", s)
	if s == "fail" {
		os.Exit(30)
	}
}
