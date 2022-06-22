package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	fmt.Printf("Przed inkrementacją, c.mux: %v, %#v, %T\n",
		c.mux, c.mux, c.mux)
	c.v[key]++
	fmt.Printf("Po inkrementacji, c.mux: %v, %#v, %T\n",
		c.mux, c.mux, c.mux)
	c.mux.Unlock()
}

// Value returns the current value of the counter for given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	fmt.Printf("Metoda Value: %v, %#v, %T\n\n",
		c.mux, c.mux, c.mux)
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}

	for i := 0; i < 3; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println("c.Value(\"somekey\"):", c.Value("somekey"))
}
