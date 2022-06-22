package main

import "fmt"

func main() {
	sample := []byte("\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")

	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Byte loop:")
	for _, elem := range sample {
		fmt.Printf("%x ", elem)
	}
	fmt.Println("\n")

	fmt.Println("Printf with: %x:")
	fmt.Printf("%x\n", sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	fmt.Println("Printf with %+q")
	fmt.Printf("%+q\n", sample)

	fmt.Printf("\n")
	fmt.Printf("sample: %v, %T\n", sample, sample)
}
