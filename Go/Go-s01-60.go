package main

import "fmt"

func main() {
	fmt.Println("Basic for-each loop (slide or array)")

	a := []string{"Foo", "Bar"}

	for i, s := range a {
		fmt.Println(i, s)
	}


	fmt.Printf("\n\nString iteration: runes or bytes\n")
	for i, ch := range "日本語" {
		fmt.Printf("%#U starts at byte position %d\n", ch, i)
	}

	fmt.Printf("\n\nLooping over individual bytes\n")

	const s = "日本語"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}


	fmt.Printf("\n\nMap iteration: keys and values\n")

	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	fmt.Println("First loop over map m")

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Second loop over map m")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Printf("\n\nChannel iteration\n")

	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()

	for n := range ch {
		fmt.Println(n)
	}
}
