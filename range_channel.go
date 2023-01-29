package main

import (
	"fmt"
	"runtime"
)

func rangeChannelMain() {
runtime.GOMAXPROCS(4)
messages := make(chan string)
go sendMessage(messages)
printMessageRange(messages)
}

func sendMessage(ch chan<- string) { 
	for i := 0; i < 20; i++ { 
		ch <- fmt.Sprintf("data %d", i) 
		} 
		
	close(ch)
 }

func printMessageRange(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}