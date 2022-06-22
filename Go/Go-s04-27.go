package main

import "fmt"

const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

func main() {
	fmt.Printf("sample: %s, %v, %#v, %T\n", sample, sample, sample,
		sample)

	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")

	fmt.Printf("%x\n", sample)
	fmt.Printf("% x\n", sample)
	fmt.Printf("%q\n", sample)
	fmt.Printf("%+q\n", sample)
}
