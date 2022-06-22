package main

import "fmt"

func WhiteSpace(c rune) bool {
	switch c {
	case ' ', '\t', '\n', '\f', '\r':
		return true
	}
	return false
}

func main() {
	fmt.Printf("WhiteSpace(%c) = %v\n", ' ', WhiteSpace(' '))
	fmt.Printf("WhiteSpace(%c) = %v\n", '\t', WhiteSpace('\t'))
	fmt.Printf("WhiteSpace(%c) = %v\n", '\f', WhiteSpace('\f'))
	fmt.Printf("WhiteSpace(%c) = %v\n", '\r', WhiteSpace('\r'))
}
