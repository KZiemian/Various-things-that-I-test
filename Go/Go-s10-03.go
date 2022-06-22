package main

import (
	"fmt"
)

func main() {
	digits := "0123456789"

	for _, v := range digits {
		if v == '0' {
			fmt.Printf("v: %v, %v == '0': true\n", v, v)
		} else if v == '1' {
			fmt.Printf("v: %v, %v == '1': true\n", v, v)
		} else if v == '2' {
			fmt.Printf("v: %v, %v == '2': true\n", v, v)
		} else if v == '3' {
			fmt.Printf("v: %v, %v == '3': true\n", v, v)
		} else if v == '4' {
			fmt.Printf("v: %v, %v == '4': true\n", v, v)
		} else if v == '5' {
			fmt.Printf("v: %v, %v == '5': true\n", v, v)
		} else if v == '6' {
			fmt.Printf("v: %v, %v == '6': true\n", v, v)
		} else if v == '7' {
			fmt.Printf("v: %v, %v == '7': true\n", v, v)
		} else if v == '8' {
			fmt.Printf("v: %v, %v == '8': true\n", v, v)
		} else if v == '9' {
			fmt.Printf("v: %v, %v == '9': true\n", v, v)
		}
	}
}
