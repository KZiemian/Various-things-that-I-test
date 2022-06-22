package main

import "fmt"

type someInterface[3] interface {
	SomeMethod()
}

type someArrayWithThreeInts [3]int

type someOtherInterface[5] interface {
	OtherMethod()
}

type someInt int

func (sI someInt) SomeMethod() {
	fmt.Println("someInt")
}

type someThreeSomeInts [3]someInt

type someOtherInt int

func (sOI someOtherInt) SomeMethod() {
	fmt.Println("someOtherInt")
}

type someAnotherInt int

func (sAI someAnotherInt) OtherMethod() {
	fmt.Println("someAnotherInt")
}

func main() {
	var varSI someInterface
	// var varSOI someOtherInterface = someAnotherInt(2)

	fmt.Printf("varSI value: %v\n", varSI)
	fmt.Printf("varSI type:  %T\n", varSI)


	varSI = someThreeSomeInts([3]someInt{someInt(1), someInt(2),
		someInt(3)})

	fmt.Printf("varSI value: %v\n", varSI)
	fmt.Printf("varSI type:  %T\n", varSI)

	// arrayVar := someArrayWithThreeInts([3]int{1, 2, 3})

	// fmt.Printf("arrayVar value: %v\n", arrayVar)
	// fmt.Printf("arrayVar type:  %T\n", arrayVar)


	// varSI = arrayVar
	// fmt.Printf("varSI value: %v\n", varSI)

	// fmt.Printf("varSOI: %T, %v\n", varSOI, varSOI)

	// varSI = someOtherInt(5)

	// fmt.Printf("varSI: %T, %v\n", varSI, varSI)
}
