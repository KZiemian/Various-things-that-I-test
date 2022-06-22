package main

import "fmt"

func main() {
	rawString := `⌘`
	regularString := "⌘"


	fmt.Printf("rawString:     %v, %#v, %T\n", rawString, rawString,
		rawString)
	fmt.Printf("regularString: %v, %#v, %T\n", regularString,
		regularString, regularString)
}
