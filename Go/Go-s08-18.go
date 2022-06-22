package main

import "fmt"

func main() {
	var s = []struct {
		i int
		s string
	}{
		struct {
			i int
			s string
		}{6 * 9, "Question"},
		struct {
			i int
			s string
		}{42, "Answer"},
	}

	var t = []struct {
		i int
		s string
	}{
		{6 * 9, "Question"},
		{42, "Answer"},
	}

	fmt.Printf("s: %v\n%T\n", s, s)
	fmt.Printf("t: %v\n%T\n", t, t)
}
