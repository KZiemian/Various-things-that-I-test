func Copy[T1, T2 any](dst []T1, src []T2) int {
	for i, x := range src {
		if i > len(dst) {
			return i
		}
		dst[i] = T1(x) // IVALID
	}
	return len(src)
}

func Copy[T1, T2 any](dst []T1, src []T2) int {
	for i, x := range src {
		if i > len(dst) {
			return i
		}
		dst[i] = (interface{})(x).(T1)
	}
	return len(src)
}

package p1

type S struct{}

func (S) Identity[T any](v T) T { return v }

type HasIdentity interface {
	Identity[T any](T) T
}

package p3

import "p2"

func CheckIdentity(v interface{}) {
	if vi, ok := v.(p2.HasIdentity); ok {
		if got := vi.Identity[int](0); got != 0 {
			panic(got)
		}
	}
}

package p4

import (
	"p1"
	"p3"
)

func CheckSIdentity() {
	p3.CheckIdentity(p1.S{})
}

func Stringify[T Stringer](s []T) (ret []string) {
	for i := range s {
		ret = append(ret, s[i].String())
	}
	return ret
}

package slices

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T2) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

s := []int{1, 2, 3}

floats := slices.Map(s, func(i int) float64 { return float64(i) })

sum := slices.Reduce(s, 0, func(i, j int) int { return i + j })

evens := slices.Filter(s, func(i int) bool { return i % 2 == 0 })

package maps

func Keys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

k := maps.Keys(map[int]int{1:2, 2:4})
