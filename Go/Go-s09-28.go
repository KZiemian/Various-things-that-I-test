package main

import "fmt"

func main() {
	fmt.Println("0b101:", 0b101)
	fmt.Println("0o10:", 0o10)
	fmt.Println("0O10:", 0O10)
	fmt.Println("010:", 010)
	fmt.Println("0b111i:", 0b111i)

	fmt.Println("0xa.cp3:", 0xa.cp3)

	var s int64 = 2
	fmt.Println("1 << s:", 1 << s)

	fmt.Println("1_000_000:", 1_000_000)
	fmt.Println("0b_1_0000:", 0b_1_0000)
	fmt.Println("02_02_2020:", 02_02_2020)
	fmt.Println("0x_FF_AA_EE:", 0x_FF_AA_EE)
}
