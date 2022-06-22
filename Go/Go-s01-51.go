package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("err: %v,\n %#v,\n %T\n\n",
			err, err, err)

		// Kod ponieżej nie działa, bo err jak wynik
		// zwrócony przez funkcję run() ma typ error,
		// który jest interfejsem.
		// fmt.Println(err.When)
		// fmt.Printf("err.When: %v,\n %#v,\n %T\n",
		// 	err.When, err.When, err.When)
	}
}
