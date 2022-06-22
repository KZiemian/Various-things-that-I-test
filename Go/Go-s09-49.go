package main

import "fmt"

func main() {
	fmt.Printf("%%d: %d\n", 10)
	fmt.Printf("%%+d: %+d\n", 10)
	fmt.Printf("|%%4d|: |%4d|\n", 10)
	fmt.Printf("|%%-4d|: |%-4d|\n", 10)
	fmt.Printf("|%%04d|: |%04d|\n", 10)
	fmt.Printf("%%b: %b\n", 10)
	fmt.Printf("%%o: %o\n", 10)
	fmt.Printf("%%x: %x\n", 10)
	fmt.Printf("%%X: %X\n", 10)
	fmt.Printf("%%#x: %#x\n", 10)
	fmt.Printf("%%#X: %#X\n", 10)

	fmt.Printf("123.456, %%e: %e\n", 123.456)
	fmt.Printf("123.456, %%f: %f\n", 123.456)
	fmt.Printf("123.456, %%.2f: %.2f\n", 123.456)
	fmt.Printf("123.456, |%%8.2f|: |%8.2f|\n", 123.456)
	fmt.Printf("123.456, %%g: %g\n", 123.456)


	i := 0
	p := &i

	fmt.Printf("p, %%v: %v\n", p)
	fmt.Printf("p, %%#v: %#v\n", p)
	fmt.Printf("p, %%T: %T\n", p)
	fmt.Printf("p, %%p: %p\n", p)

	fmt.Printf("true, %%v: %v\n", true)
	fmt.Printf("true, %%#v: %#v\n", true)
	fmt.Printf("true, %%T: %T\n", true)
	fmt.Printf("true, %%t: %t\n", true)

	fmt.Printf("Unicode letter b (number 98), %%v: %v\n", 98)
	fmt.Printf("Unicode letter b (number 98), %%#v: %#v\n", 98)
	fmt.Printf("Unicode letter b (number 98), %%T: %T\n", 98)
	fmt.Printf("Unicode letter b (number 98), %%c: %c\n", 98)
	fmt.Printf("Unicode letter b (number 98), %%q: %q\n", 98)
	fmt.Printf("Unicode letter b (number 98), %%U: %U\n", 98)
	fmt.Printf("Unicode letter b (number 98), %%#U: %#U\n", 98)

	fmt.Printf("\"café\", %%s: %s\n", "café")
	fmt.Printf("\"café\", |%%6s|: |%6s|\n", "café")
	fmt.Printf("\"café\", %%x: %x\n", "café")
	fmt.Printf("\"café\", %% x: % x\n", "café")

	fmt.Printf("[]int{0, 1}, %%v: %v\n", []int{0, 1})
	fmt.Printf("[]int{0, 1}, %%#v: %#v\n", []int{0, 1})
	fmt.Printf("[]int{0, 1}, %%T: %T\n", []int{0, 1})

	fmt.Printf("[]int64{0, 1}, %%v: %v\n", []int64{0, 1})
	fmt.Printf("[]int64{0, 1}, %%#v: %#v\n", []int64{0, 1})
	fmt.Printf("[]int64{0, 1}, %%T: %T\n", []int64{0, 1})

	fmt.Println("caf\u00e9")
}
