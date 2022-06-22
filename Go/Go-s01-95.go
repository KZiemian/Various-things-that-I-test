package main

import "fmt"

func main() {
	b := []byte("ABC€")

	fmt.Println("String to byte slice")
	fmt.Println(b)
	fmt.Printf("b: %v, %#v, %T\n\n", b, b, b)


	s := string([]byte{65, 66, 67, 226, 130, 172})

	fmt.Println("Byte slice to string")
	fmt.Println(s)
	fmt.Printf("s: %v, %#v, %T\n", s, s, s)
}
