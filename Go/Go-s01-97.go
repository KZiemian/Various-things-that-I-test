package main

import "fmt"

func main() {
	r := []rune("ABC€")

	fmt.Println(r)
	fmt.Printf("%U\n", r)

	s := string([]rune{'\u0041', '\u0042', '\u0043', '\u20AC', -1})
	fmt.Println(s)
	fmt.Printf("%v, %#v, %T\n", s, s, s)
}
