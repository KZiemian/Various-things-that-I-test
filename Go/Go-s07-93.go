package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := regexp.MustCompile("a*").Split("abaabaccadaaae", 5)
	fmt.Printf("%q\n", s)
}
