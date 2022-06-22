package main

import "fmt"

type channelWithDescription struct {
	channelName string
	channelField chan int
}

func main() {
	varStructChannel := channelWithDescription{"Int Channel Dave", make(chan int, 3)}

	varStructChannel.channelField <- 0
	varStructChannel.channelField <- 1
	varStructChannel.channelField <- 2

	fmt.Println("varStructChannel.channelName:",
		varStructChannel.channelName)
	fmt.Printf("varStructChannel.channelField type: %T\n",
		varStructChannel.channelField)
	fmt.Println("<-varStructChannel.channelField, 1:",
		<-varStructChannel.channelField)
	fmt.Println("<-varStructChannel.channelField, 2:",
		<-varStructChannel.channelField)
	fmt.Println("<-varStructChannel.channelField, 3:",
		<-varStructChannel.channelField)
}
