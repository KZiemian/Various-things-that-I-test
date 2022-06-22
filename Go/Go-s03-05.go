package main

import "fmt"

type channelStruct struct {
	Name string
	ChannelField chan int
}

func main() {
	varChannelStruct := channelStruct{"Channel Dave",
		make(chan int)}


	go func() {
		varChannelStruct.ChannelField <- 0
		varChannelStruct.ChannelField <- 1
		varChannelStruct.ChannelField <- 2
	}()

	fmt.Println("varChannelStruct.Name:", varChannelStruct.Name)
	fmt.Println("<-varChannelStruct.ChannelField, 1:",
		<-varChannelStruct.ChannelField)
	fmt.Println("<-varChannelStruct.ChannelField, 2:",
		<-varChannelStruct.ChannelField)
	fmt.Println("<-varChannelStruct.ChannelField, 3:",
		<-varChannelStruct.ChannelField)
}
