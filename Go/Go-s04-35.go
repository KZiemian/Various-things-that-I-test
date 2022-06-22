package main

import "fmt"

func main() {
	const nihongo = "日本語"

	fmt.Printf("nihongo: %v, %#v, %T\n", nihongo, nihongo, nihongo)
	fmt.Printf("nihongo: %s\n", nihongo)
}
