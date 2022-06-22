type Point []int32

func ScaleAndPrint(p Point) {
	r := Scale(p, 2) // calls Scale[Point, int32]
	// ...
}

func Scale[S ~[]E, E constraints.Integer](s S, sc E) S {
	// ...
}

type Tree[T any] struct {
	cmp  func(T, T) int
	root *node[t]
}

type node[T any] struct {
	left, right *node[T]
	data        T
}

func (bt *Tree[T]) find(val T) **node[T] {
	pl := &bt.root
	for *pl != nil {
		switch cmp := bt.cmp(val, (*pl).data); {
		case cmp < 0:
			pl = &(*pl).left
		case cmp > 0:
			pl = &(*pl).right
		default:
			return pl
		}
	}
	return pl
}

type SliceFn[T any] struct {
	s   []T
	cmp func(T, T) bool
}

func (s SliceFn[T]) Len() int { return len(s.s) }

func (s SliceFn[T]) Swap(i, j int) {
	s.s[i], s.s[j] = s.s[j], s.s[i]
}

func (s SliceFn[T]) Less(i, j int) bool {
	return s.cmp(s.s[i], s.s[j])
}

func SortFn[T any](s []T, cmp func(T, T) bool) {
	sort.Sort(SliceFn[T]{s, cmp})
}

package main

import (
	"fmt"
	"errors"
	"unicode/utf8"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, _ := Reverse(input)
	dobuleRev, _ := Reverse(rev)

	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)
}

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r) nil
}

package main

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, dobuleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q",
				rev)
		}
	})
}
