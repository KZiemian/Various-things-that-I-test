package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var n uint = 0

	fmt.Printf("%d (%b) has %d bits set to one\n", n, n,
		bits.OnesCount(n))
	fmt.Printf("%d reversed is %d\n", n, bits.Reverse(n))
	fmt.Printf("%d can be encoded in %d bits\n", n, bits.Len(n))

	n = 1
	fmt.Printf("%d (%b) has %d bits set to one\n", n, n,
		bits.OnesCount(n))
	fmt.Printf("%d reversed is %d\n", n, bits.Reverse(n))
	fmt.Printf("%d can be encoded in %d bits\n", n, bits.Len(n))

	n = 21
	fmt.Printf("%d (%b) has %d bits set to one\n", n, n,
		bits.OnesCount(n))
	fmt.Printf("%d reversed is %d\n", n, bits.Reverse(n))
	fmt.Printf("%d can be encoded in %d bits\n", n, bits.Len(n))
}
