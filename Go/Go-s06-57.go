package main

import "fmt"

func main() {
	const invalidUTF8 = "\xbd\xb2\xbc"

	for index, value := range invalidUTF8 {
		fmt.Printf("%#U starts at byte position %d\n", value, index)
	}
}
