package main

import "fmt"

func main() {
	const badUTF8 = "\xbd\xbd\xb2\xb2日本語"

	for index, runeValue := range badUTF8 {
		fmt.Printf("%#U starts at byte position %d\n", runeValue,
			index)
	}

	fmt.Printf("len(badUTF8): %v\n", len(badUTF8))
}
