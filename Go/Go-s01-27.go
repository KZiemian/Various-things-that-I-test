package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Printf("a: %v, %#v, %T\n\n", a, a, a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Printf("primes: %v, %#v, %T\n\n", primes, primes,
		primes)

	var primesTwo [6]int = [6]int{2, 3, 5, 7, 11, 13}
	fmt.Printf("primesTwo: %v, %#v, %T\n", primesTwo,
		primesTwo, primesTwo)
}
