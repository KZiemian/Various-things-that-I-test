package main

import (
	"fmt"
)

func main() {
	var varInt   int   = 0
	var varInt8  int8  = 0
	var varInt16 int16 = 0
	var varInt32 int32 = 0

	var varUint   uint   = 0
	var varUint8  uint8  = 0
	var varUint16 uint16 = 0
	var varUint32 uint32 = 0
	var varUint64 uint64 = 0

	// var varUintpointer uintptr = uintptr(&varInt)

	var varByte byte = 0

	var varRune rune = 0

	var varFloat32 float32 = 0.0

	var varFloat64 float64 = 0.0

	var varComplex64 complex64 = 0
	var varComplex128 complex128 = 0

	fmt.Printf("  varInt: %v, %#v, %T\n", varInt, varInt,
		varInt)
	fmt.Printf(" varInt8: %v, %#v, %T\n", varInt8, varInt8,
		varInt8)
	fmt.Printf("varInt16: %v, %#v, %T\n", varInt16, varInt16,
		varInt16)
	fmt.Printf("varInt32: %v, %#v, %T\n\n\n", varInt32,
		varInt32, varInt32)

	fmt.Printf("  varUint: %v, %#v, %T\n", varUint, varUint,
		varUint)
	fmt.Printf(" varUint8: %v, %#v, %T\n", varUint8,
		varUint8, varUint8)
	fmt.Printf("varUint16: %v, %#v, %T\n", varUint16,
		varUint16, varUint16)
	fmt.Printf("varUint32: %v, %#v, %T\n", varUint32,
		varUint32, varUint32)
	fmt.Printf("varUint64: %v, %#v, %T\n\n\n", varUint64,
		varUint64, varUint64)

	// fmt.Printf("varUintptr: %v, %#v, %T\n\n\n", varUintptr,
	// 	varUintptr, varUintptr)

	fmt.Printf("varByte: %v, %#v, %T\n\n\n", varByte, varByte,
		varByte)

	fmt.Printf("varRune: %v, %#v, %T\n\n\n", varRune, varRune,
			varRune)

	fmt.Printf("varFloat32: %v, %#v, %T\n", varFloat32,
		varFloat32, varFloat32)

	fmt.Printf("varFloat64: %v, %#v, %T\n\n\n", varFloat64,
		varFloat64, varFloat64)

	fmt.Printf("varComplex64: %v, %#v, %T\n",
		varComplex64, varComplex64, varComplex64)
	fmt.Printf("varComplex128: %v, %#v, %T\n", varComplex128,
		varComplex128, varComplex128)
}
