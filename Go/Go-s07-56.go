package main

import (
	"fmt"
	""
)

var input = []string{"Sir", "Martin", "Hairer", "KBE", "FRS", "(born", "14",
	"November", "1975[1])", "is", "an", "Austrian-British",
	"mathematician", "working", "in", "the", "field", "of", "stochastic",
	"analysis", "in", "particular", "stochastic", "partial",
	"differential", "equations.", "He", "is", "Professor", "of",
	"Mathematics", "at", "Imperial", "College", "London,", "having",
	"previously", "held", "appointments", "at", "the", "University",
	"of", "Warwick", "and", "the", "Courant", "Institute", "of", "New",
	"York", "University."}

func copyList1(in []string) []string {
	var out []string
	for _, s := range in {
		out = append(out, s)
	}
	return our
}

func copyList2(in []string) []string {
	out := make([]string, len(in))
	for i, s := range in {
		out[i] = s
	}
	return out
}

func BenchmarkCopyList1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		copyList1(input)
	}
}

func BenchmarkCopyList2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		copyList2(input)
	}
}

func main() {

}
