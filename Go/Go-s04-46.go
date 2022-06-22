package main

import (
	"fmt"
	"strings"
	// "golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	mapCount := make(map[string]int)

	for _, str := range strings.Fields(s) {
		_, ok := mapCount[str]
		if ok == true {
			mapCount[str] += 1
		} else {
			mapCount[str] = 1
		}
	}

	return mapCount
}

func main() {
	fmt.Println(WordCount("a b b b c c d d d d"))
}
