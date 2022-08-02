package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Simple Shell")
	fmt.Println("----------------------")

	for i := 0; i < 10; i++ {
		fmt.Print("->")
		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\r\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourslef")
		}
	}
}
