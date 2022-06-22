package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

func (l List[T]) ShowList(number int) {
	fmt.Printf("Element of number %v: %v\n", number, l.val)

	if l.next != nil {
		(l.next).ShowList(number + 1)
	} else {
		fmt.Println("End of the list.")
	}
}

func main() {
	var l4 List[int] = List[int]{nil, 23}
	var l3 List[int] = List[int]{&l4, 17}
	var l2 List[int] = List[int]{&l3, 13}
	var l1 List[int] = List[int]{&l2, 7}
	var l0 List[int] = List[int]{&l1, 1}

	fmt.Printf("l0.next: %v, %T\n", l0.next, l0.next)
	fmt.Printf("l1.next: %v, %T\n", l1.next, l1.next)
	fmt.Printf("l2.next: %v, %T\n", l2.next, l2.next)
	fmt.Printf("l3.next: %v, %T\n", l3.next, l3.next)
	fmt.Printf("l4.next: %v, %T\n", l4.next, l4.next)
	fmt.Printf("\n\n\n")

	l0.ShowList(0)
}
