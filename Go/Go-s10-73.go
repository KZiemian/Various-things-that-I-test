package main

import "fmt"

func main() {
	varInt := 42
	fmt.Println("42:", varInt)
	varInt = 4_2
	fmt.Println("4_2:", varInt)
	varInt = 0600
	fmt.Println("0600:", varInt)
	varInt = 0_600
	fmt.Println("0_600:", varInt)
	varInt = 0o600
	fmt.Println("0o600:", varInt)
	varInt = 0O600
	fmt.Println("0O600:", varInt)
	varInt = 0xBadFace
	fmt.Println("0xBadFace:", varInt)
	varInt = 0xBad_Face
	fmt.Println("0xBad_Face:", varInt)
	varInt = 0x_67_7a_2f_cc_40_c6
	fmt.Println("0x_67_7a_2f_cc_40_c6:", varInt)
}
