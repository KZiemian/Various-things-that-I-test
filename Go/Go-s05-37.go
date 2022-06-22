package main

import (
	"fmt"
	"time"
)

const millisecond = time.Second/1e3

func main() {
	bigBufferWithHeader := make([]byte, 512+1e6)

	fmt.Printf("%v, %T\n", millisecond, millisecond)
	fmt.Printf("%T, %v, %v\n", bigBufferWithHeader,
		len(bigBufferWithHeader), cap(bigBufferWithHeader))
}
