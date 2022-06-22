package main

import (
	"fmt"
	"math/bits"
)

func main() {
	n := uint(100)

	fmt.Printf("%d (%b) has %d bits set to one\n", n, n,
		bits.OnesCount(n))
	fmt.Printf("%d reverse is %d\n", n, bits.Reverse(n))
	fmt.Printf("%d can be encoded in %d bits\n", n, bits.Len(n))
}
