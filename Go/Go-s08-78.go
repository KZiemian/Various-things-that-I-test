package sets

typ Set[T comparable] map[T]struct{}

func Make[T comparable]() Set(T) {
	return make(Set[T])
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Delete(v T) {
	delete(s, v)
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	retunr ok
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Iterate(f func(T)) {
	for v := range s {
		f(v)
	}
}

s := sets.Make[int]()

s.Add(1)

if s.Constains(2) { panic("unexpected 2") }

type Ordered interface {
	~int | ~int8 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

type orderedSlice[T Ordered] []T

func (s orderedSlice[T]) Len() int { return len(s) }
func (s orderedSlice[T]) Less(i, j int) bool { return s[i] < s[j] }

func OrderedSlice[T Ordered](s []T) {
	sort.Sort(orderedSlice[T](s))
}

s1 := []int32{3, 5, 2}
sort.OrderedSlice(s1)

s2 := []string{"a", "b", "c"}
sort.OrderedSlice(s2)

type sliceFn[T any] struct {
	s   []T
	cmp func(T, T) bool
}

func (s sliceFn[T]) Len() int           { return len(s.s) }
func (s sliceFn[T]) Less(i, j int) bool { return s.cmp(s.s[i], s.s[j]) }
func (s sliceFn[T]) Swap(i, j int)      { s.s[i], s.s[j] = s.s[j], s.s[i] }

func SliceFn[T any](s []T, cmp func(T, T) bool) {
	sort.Sort(sliceFn[T]{s, cmp})
}

var s []*Person
// ...
sort.SliceFn(s, func(p1, p2 *Person) bool { return p1.Name < p2.Name })

package chans

import "runtime"

func Drain[T any](c <-chan T) {
	for range c {
	}
}

func Merge[T any](c1, c2 <-chan T) <-chan T {
	r := make(chan T)
	go func(c1, c2 <-chan T, r chan<- T) {
		defer close(r)
		for c1 != nil || c2 != nil {
			select {
			case v1, ok := <-c1:
				if ok {
					r <- v1
				} else {
					c1 = nil
				}
			case v2, ok := <-c2:
				if ok {
					r <- v2
				} else {
					c2 = nil
				}
			}
		}
	}(c1, c2, r)
	return r
}
