package main

import "fmt"

func main() {
	stringUsr := "/usr/ken"
	slash := "/usr/ken"[0]

	fmt.Printf("stringUsr: %s, %v, %#v, %T\n", stringUsr, stringUsr,
		stringUsr, stringUsr)
	fmt.Printf("slash: %v, %#v, %T\n", slash, slash, slash)

	usr := "/usr/ken"[0:4]
	fmt.Printf("usr: %s, %v, %#v, %T\n", usr, usr, usr, usr)

	slice := []byte{65, 66, 67, 68, 69, 70}
	fmt.Printf("slice: %v, %#v, %T\n", slice, slice, slice)

	str := string(slice)
	fmt.Printf("str: %s, %v, %#v, %T\n", str, str, str, str)

	sliceStr := []byte(usr)
	fmt.Printf("sliceStr: %v, %#v, %T\n", sliceStr, sliceStr, sliceStr)
}
