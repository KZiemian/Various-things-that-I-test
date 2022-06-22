package main

import "fmt"

func main() {
	v := 'a'
	fmt.Printf("'a': %v, %T\n", v, v)
	v = 'ä'
	fmt.Printf("'ä': %v, %T\n", v, v)
	v = '本'
	fmt.Printf("'本': %v, %T\n", v, v)
	v = '\t'
	fmt.Printf("'\t': %v, %T\n", v, v)
	v = '\000'
	fmt.Printf("'\000': %v, %T\n", v, v)
	v = '\007'
	fmt.Printf("'\007': %v, %T\n", v, v)
	v = '\377'
	fmt.Printf("'\377': %v, %T\n", v, v)
	v = '\x07'
	fmt.Printf("'\x07': %v, %T\n", v, v)
	v = '\xff'
	fmt.Printf("'\xff': %v, %T\n", v, v)
	v = '\u12e4'
	fmt.Printf("'\u12e4': %v, %T\n", v, v)
	v = '\U00101234'
	fmt.Printf("'\U00101234': %v, %T\n", v, v)
	v = '\''
	fmt.Printf("': %v, %T\n", v, v)
}
