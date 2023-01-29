package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func timeoutChannelMain() {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(4)
	messages := make(chan int)

	go sendData(messages)
	retrieveData(messages)
}

func sendData(ch chan<- int) {
	for i := 0; true; i++ {
		ch <- i
		timeSleep := rand.Int()%10 + 1
		// timeSleep /= timeSleep
		time.Sleep(time.Duration(timeSleep) * time.Second)
	}
}

func retrieveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Print(`receive data "`, data, `"`, "\n")

		case <-time.After(time.Second * 5):
			fmt.Println("timeout. no activities under 5 seconds")
			break loop
		}
	}
}
