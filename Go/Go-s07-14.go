package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println("match:", match)

	match, _ = regexp.MatchString("p([a-z]+)ch", "pch")
	fmt.Println("match:", match)

	fmt.Printf("match: %v, %T\n", match, match)
}
