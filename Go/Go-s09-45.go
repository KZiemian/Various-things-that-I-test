package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	fmt.Printf("%%\n")

	fmt.Printf("%%v: %v\n", 10)
	fmt.Printf("%%#v: %#v\n", 10)
	fmt.Printf("%%+v: %+v\n", 10)
	fmt.Printf("%%T: %T\n", 10)

	v1 := Vertex{0, 1}
	fmt.Printf("%%v: %v\n", v1)
	fmt.Printf("%%#v: %#v\n", v1)
	fmt.Printf("%%+v: %+v\n", v1)
	fmt.Printf("%%T: %T\n", v1)

	v2 := true
	fmt.Printf("%%t: %t\n", v2)

	i1 := 10
	i2 := 65
	fmt.Printf("%%b: %b\n", i1)
	fmt.Printf("%%c: %c\n", i2)
	fmt.Printf("%%d: %d\n", i1)
	fmt.Printf("%%d: %d\n", i2)
	fmt.Printf("%%o: %o\n", i1)
	fmt.Printf("%%o: %o\n", i2)
	fmt.Printf("%%O: %O\n", i1)
	fmt.Printf("%%O: %O\n", i2)
	fmt.Printf("%%q: %q\n", i1)
	fmt.Printf("%%q: %q\n", i2)
	fmt.Printf("%%x: %x\n", i1)
	fmt.Printf("%%X: %X\n", i1)
	fmt.Printf("%%U: %U\n", i1)
	fmt.Printf("%%U: %U\n", i2)

	f1 := 0.5
	f2 := 2.0
	f3 := 2.5
	f4 := 64.0

	fmt.Printf("%%b: %b\n", f1)
	fmt.Printf("%%b: %b\n", f2)
	fmt.Printf("%%b: %b\n", f3)
	fmt.Printf("%%b: %b\n", f4)
	fmt.Printf("%%e: %e\n", f1)
	fmt.Printf("%%e: %e\n", f2)
	fmt.Printf("%%e: %e\n", f3)
	fmt.Printf("%%e: %e\n", f4)
	fmt.Printf("%%E: %E\n", f1)
	fmt.Printf("%%E: %E\n", f2)
	fmt.Printf("%%E: %E\n", f3)
	fmt.Printf("%%E: %E\n", f4)
	fmt.Printf("%%f: %f\n", f1)
	fmt.Printf("%%f: %f\n", f2)
	fmt.Printf("%%f: %f\n", f3)
	fmt.Printf("%%f: %f\n", f4)
	fmt.Printf("%%F: %F\n", f1)
	fmt.Printf("%%F: %F\n", f2)
	fmt.Printf("%%F: %F\n", f3)
	fmt.Printf("%%F: %F\n", f4)
	fmt.Printf("%%g: %g\n", f1)
	fmt.Printf("%%g: %g\n", f2)
	fmt.Printf("%%g: %g\n", f3)
	fmt.Printf("%%g: %g\n", f4)
	fmt.Printf("%%G: %G\n", f1)
	fmt.Printf("%%G: %G\n", f2)
	fmt.Printf("%%G: %G\n", f3)
	fmt.Printf("%%G: %G\n", f4)
	fmt.Printf("%%x: %x\n", f1)
	fmt.Printf("%%x: %x\n", f2)
	fmt.Printf("%%x: %x\n", f3)
	fmt.Printf("%%x: %x\n", f4)
	fmt.Printf("%%X: %X\n", f1)
	fmt.Printf("%%X: %X\n", f2)
	fmt.Printf("%%X: %X\n", f3)
	fmt.Printf("%%X: %X\n", f4)

	slice1 := []byte{65, 66, 67, 68, 69, 70}
	string1 := "string"
	fmt.Printf("%%s: %s\n", slice1)
	fmt.Printf("%%s: %s\n", string1)
	fmt.Printf("%%q: %q\n", slice1)
	fmt.Printf("%%q: %q\n", string1)
	fmt.Printf("%%x: %x\n", slice1)
	fmt.Printf("%%x: %x\n", string1)
	fmt.Printf("%%X: %X\n", slice1)
	fmt.Printf("%%X: %X\n", string1)

	fmt.Printf("%%p: %p\n", slice1)
	p1 := &i1
	fmt.Printf("%%p: %p\n", p1)
	fmt.Printf("%%b: %b\n", p1)
	fmt.Printf("%%d: %d\n", p1)
	fmt.Printf("%%o: %o\n", p1)
	fmt.Printf("%%x: %x\n", p1)
	fmt.Printf("%%X: %X\n", p1)

	// bool: %v -> %t
	// int, int8 etc.: %v -> %d
	// uint, uint8 etc.: %v -> %d, %#v -> %#x
	// float32, complex64, etc.
	// string: %v -> %s
	// chan: %v -> %p
	// pointer: %v -> %p
}
