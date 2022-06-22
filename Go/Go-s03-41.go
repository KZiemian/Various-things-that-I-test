package main

import "fmt"

var (
	varBool       bool       = true
	varString     string     = "This is a string."
	varInt        int        = 1
	varInt8       int8       = 1
	varInt16      int16      = 1
	varInt32      int32      = 1
	varInt64      int64      = 1
	varUint       uint       = 1
	varUint8      uint8      = 1
	varUint16     uint16     = 1
	varUint32     uint32     = 1
	varUint64     uint64     = 1
	varUintptr    uintptr
	varByte       byte       = 1
	varRune       rune       = 'A'
	varFloat32    float32    = 1.0
	varFloat64    float64    = 1.0
	varComplex64  complex64  = 1.0 + 1.0i
	varComplex128 complex128 = 1.0 + 1.0i
)

func main() {
	fmt.Printf("varBool:    %v, %#v, %T\n", varBool, varBool, varBool)
	fmt.Printf("varString:  %v, %#v, %T\n", varString, varString,
		varString)
	fmt.Printf("varInt:     %v,  %#v,   %T\n", varInt, varInt, varInt)
	fmt.Printf("varInt8:    %v,  %#v,   %T\n", varInt8, varInt8, varInt8)
	fmt.Printf("varInt16:   %v,  %#v,   %T\n", varInt16, varInt16, varInt16)
	fmt.Printf("varInt32:   %v,  %#v,   %T\n", varInt32, varInt32, varInt32)
	fmt.Printf("varInt64:   %v,  %#v,   %T\n", varInt64, varInt64, varInt64)
	fmt.Printf("varUint:    %v,  %#v, %T\n", varUint, varUint, varUint)

	fmt.Printf("varUint8:   %v,  %#v, %T\n", varUint8, varUint8, varUint8)
	fmt.Printf("varUint16:  %v,  %#v, %T\n", varUint16, varUint16,
		varUint16)
	fmt.Printf("varUint32:  %v,  %#v, %T\n", varUint32, varUint32,
		varUint32)
	fmt.Printf("varByte:    %v,  %#v, %T\n", varByte, varByte, varByte)
	fmt.Printf("varRune:    %v, %#v,  %T\n", varRune, varRune, varRune)
	fmt.Printf("varFloat32: %v,  %#v,   %T\n", varFloat32, varFloat32,
		varFloat32)
	fmt.Printf("varFloat64: %v,  %#v,   %T\n", varFloat64, varFloat64,
		varFloat64)
	fmt.Printf("varComplex64:  %v, %#v, %T\n", varComplex64, varComplex64,
		varComplex64)
	fmt.Printf("varComplex128: %v, %#v, %T\n", varComplex128,
		varComplex128, varComplex128)
}
