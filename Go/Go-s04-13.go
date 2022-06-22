package main

import "fmt"

type path []byte

func (p *path) ToUpper() {
	for i, b := range *p {
		if 'a' <= b && b <= 'z' {
			(*p)[i] = b + 'A' - 'a'
		}
	}
}

func main() {
	pathName := path("/usr/bin/tso")
	pathName.ToUpper()
	fmt.Printf("%s\n", pathName)
}
