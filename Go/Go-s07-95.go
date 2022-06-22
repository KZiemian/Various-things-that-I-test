package main

import (
	"fmt"
	"regexp"
)

func main() {
	zp := regexp.MustCompile(`z+`)
	fmt.Println(zp.Split("pizza", -1))
	fmt.Println(zp.String())
}
