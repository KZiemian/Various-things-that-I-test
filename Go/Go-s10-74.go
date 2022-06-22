package main

import "fmt"

func main() {
	varFloat := 0.
	// fmt.Printf("varFloat: %T\n", varFloat)
	fmt.Println("0.:", varFloat)
	varFloat = 72.40
	fmt.Println("72.40:", varFloat)
	varFloat = 072.40
	fmt.Println("072.40:", varFloat)
	varFloat = 2.71828
	fmt.Println("2.71828:", varFloat)
	varFloat = 1.e+0
	fmt.Println("1.e+0:", varFloat)
	varFloat = .12345E+5
	fmt.Println(".12345E+5:", varFloat)
	varFloat = 6.67428e-11
	fmt.Println("6.67428e-11:", varFloat)
	varFloat = 1E6
	fmt.Println("1E6:", varFloat)
	varFloat = .25
	fmt.Println(".25:", varFloat)
	varFloat = .12345E+5
	fmt.Println(".12345E+5:", varFloat)
	varFloat = 1_5.
	fmt.Println("1_5.:", varFloat)
	varFloat = 0.15e+0_2
	fmt.Println("0.15e+0_2:", varFloat)
	fmt.Println("")

	varFloat = 0x1p-2
	fmt.Println("0x1p-2:", varFloat)
	varFloat = 0x2p10
	fmt.Println("0x2.p10:", varFloat)
	varFloat = 0x1.Fp+0
	fmt.Println("0x1.Fp+0:", varFloat)
	varFloat = 0X.8p-0
	fmt.Println("0X.8p-0:", varFloat)
	varFloat = 0X_1FFFP-16
	fmt.Println("0X_1FFFP-16:", varFloat)
	varFloat = 0x15e-2
	fmt.Println("0x15e-2:", varFloat)
}
