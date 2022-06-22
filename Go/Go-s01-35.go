package main

import "fmt"

func main() {
	sum := 0

	for i := 1; i < 5; i++ {
		sum += i
	}

	fmt.Println("sum:", sum)
	fmt.Printf("sum: %v, %#v, %T\n\n", sum, sum, sum)


	fmt.Println("Ponownie for i := 1; i < 5; i++")

	sum = 0
	for i := 1; i < 5; i++ {
		sum += i
		fmt.Printf("i: %v, %#v, %T\n", i, i, i)
		fmt.Printf("sum: %v, %#v, %T\n", sum, sum, sum)
	}
	fmt.Printf("\nsum: %v, %#v, %T\n\n", sum, sum, sum)


	fmt.Println("for n < 5")

	n := 1

	for n < 5 {
		n *= 2
	}
	fmt.Println("n:", n)
	fmt.Printf("n: %v, %#v, %T\n\n", n, n, n)

	fmt.Println("Ponownie for n < 5")

	n = 1
	for n < 5 {
		n *= 2
		fmt.Printf("n: %v, %#v, %T\n", n, n, n)
	}
	fmt.Printf("n: %v, %#v, %T\n", n, n, n)


	strings := []string{"hello", "world"}
	for i, s := range strings {
		fmt.Println(i, s)
	}


	fmt.Printf("\n\nExit a loop\n")

	sum = 0
	for i := 1; i < 5; i++ {
		if i % 2 != 0 {
			continue
		}
		sum += i
	}
	fmt.Println("sum:", sum)
	fmt.Printf("sum: %v, %#v, %T\n", sum, sum, sum)


	fmt.Printf("\n\nExit a loop, again\n")

	sum = 0
	for i := 1; i < 5; i++ {
		if i % 2 != 0 {
			continue
		}
		sum += i
		fmt.Printf("i: %v, %#v, %T\n", i, i, i)
		fmt.Printf("sum: %v, %#v, %T\n", sum, sum, sum)
	}
	fmt.Printf("\nsum: %v, %#v, %T\n", sum, sum, sum)
}
