package main

import "fmt"

func main() {
	var u uint
	const v = -1
	u = uint(v)

	fmt.Println(u)
}
