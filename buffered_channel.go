package main

import (
	"fmt"
	"runtime"
)

func bufferedChannelMain() {
	runtime.GOMAXPROCS(4)

	messsages := make(chan int, 3)

	go func() {
		for {
			i := <-messsages
			fmt.Println("received data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		messsages <- i
	}
}
