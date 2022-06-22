package main

import (
	"fmt"
	"sort"
)

func main() {
	var m map[int]string = map[int]string{1: "one", 99: "nine nie",
		3: "three", 10: "one zero", 22: "two two"}
	var keys []int

	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}
}
