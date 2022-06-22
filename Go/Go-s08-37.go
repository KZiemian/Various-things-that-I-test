package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Printf("log.Ldate: %v, %T\n", log.Ldate, log.Ldate)
	fmt.Printf("log.Ltime: %v, %T\n", log.Ltime, log.Ltime)
	fmt.Printf("log.Lmicroseconds: %v, %T\n", log.Lmicroseconds,
		log.Lmicroseconds)
	fmt.Printf("log.Llongfile: %v, %T\n", log.Llongfile, log.Llongfile)
	fmt.Printf("log.Lshortfile: %v, %T\n", log.Lshortfile, log.Lshortfile)
	fmt.Printf("log.LUTC: %v, %T\n", log.LUTC, log.LUTC)
	fmt.Printf("log.Lmsgprefix: %v, %T\n", log.Lmsgprefix, log.Lmsgprefix)
	fmt.Printf("log.LstdFlags: %v, %T\n", log.LstdFlags, log.LstdFlags)
}
