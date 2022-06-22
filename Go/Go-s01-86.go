package main

import (
	"fmt"
	"time"
)

func main() {
	timeNow := time.Now()
	timeWeekday := time.Now().Weekday()

	fmt.Printf("timeNow: %v,\n%#v,\n%T\n\n", timeNow, timeNow, timeNow)
	fmt.Printf("timeWeekday: %v,\n%#v,\n%T\n\n", timeWeekday,
		timeWeekday, timeWeekday)
}
