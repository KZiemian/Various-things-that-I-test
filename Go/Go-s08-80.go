package slice

func Append[T any](s []T, t ...T) []T {
	lens := len(s)
	tot := lens + len(t)
	if tot < 0 {
		panic("Append: cap out of range")
	}
	if tot > cap(s) {
		news := make([]T, tot, tot + tot/2)
		copy(news, s)
		s = news
	}
	s = s[:tot]
	copy(s[lens:], t)
	return s
}

func Copy[T any](s, t []T) int {
	i := 0
	for ; i < len(s) && i < len(t); i++ {
		s[i] = t[i]
	}
	return i
}

s := slices.Append([]int{1, 2, 3}, 4, 5, 6)
slices.Copy(s[3:], []int{7, 8, 9})

package metrics

import "sync"

type Metric1[T comparable] struct {
	mu sync.Mutex
	m  map[T]int
}

func (m *Metric1[T]) Add(v T) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.m == nil {
		m.m = make(map[T]int)
	}
	m.m[v]++
}

type key2[T1, T2 comparable] struct {
	f1 T1
	f2 T2
}

type Metric2[T1, T2 comparable] struct {
	mu sync.Mutex
	m  map[key2[T1, T2]]int
}

func (m *Metric2[T1, T2]) Add(v1 T1, v2 T2) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.m == nil {
		m.m = make(map[key2[T1, T2]]int)
	}
	m.m[key2[T1, T2]{v1, v2}]++
}

type key3[T1, T2, T3 comparable] struct {
	f1 T1
	f2 T2
	f3 T3
}

type Metric3[T1, T2, T3 comparable] struct {
	mu sync.Mutex
	m  map[key3[T1, T2, T3]]int
}

func (m *Metric3[T1, T2, T3]) Add(v1 T1, v2 T2, v3 T3) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.m == nil {
		m.m = make(map[key3[T1, T2, T3]]int)
	}
	m.m[key3[T1, T2, T3]{v1, v2, v3}]++
}

import "metrics"

var m = metrics.Metric2[string, int]{}

func F(s string, i int) {
	m.Add(s, i)
}

package lists

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

type Iterator[T any] struct {
	next **element[T]
}

func (lst *List[T]) Range() *Iterator[T] {
	return Iterator[T]{next: &lst.head}
}

func (it *Iterator[T]) Next() bool {
	if *it.next == nil {
		return false
	}
	it.next = &(*it.next).next
	return true
}

func (it *Iterator[T]) Val() (T, bool) {
	if *it.next == nil {
		var zero T
		return zero, false
	}
	return (*it.next).val, true
}

func Transform[T1, T2 any](lst *List[T1], f func(T1) T2) *List[T2] {
	ret := &List[T2]{}
	it := lst.Range()
	for {
		if v, ok := it.Val(); ok {
			ret.Push(f(v))
		}
		if !it.Next() {
			break
		}
	}
	return ret
}
