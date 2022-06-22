package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`a(x*)b(y|z)c`)
	fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbc-"))
	fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))
}
