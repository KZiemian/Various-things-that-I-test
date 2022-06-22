package main

import "fmt"

func main() {
Loop:
	for _, ch := range "a b\nc" {
		switch ch {
		case ' ': // skip space
			break
		case '\n':
			break Loop
		default:
			fmt.Printf("%c\n", ch)
		}
	}
}
