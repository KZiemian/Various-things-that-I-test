package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		i int
		b bool
		s string
	)
	r := strings.NewReader("5 true gophers")
	n, err := fmt.Fscan(r, &i, &b, &s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fscan: %v\n", err)
	}
	fmt.Println(i, b, s)
	fmt.Println(n)
}
