package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

var input = `
{
"created_at": "Thu May 31 00:00:01 +0000 2012"
}
`
func main() {
	var val map[string]time.Time

	if err := json.Unmarshal([]byte(input), &val); err != nil {
		panic(err)
	}

	fmt.Println(val)
	for k, v := range val {
		fmt.Println(k, reflect.TypeOf(v))
	}
}
