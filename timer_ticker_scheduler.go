package main

import (
	"fmt"
	"os"
	"time"
)

func timerMain() {
	timerWatcherMain()
	done := make(chan bool)
	ticker := time.NewTicker(time.Second)

	go func() {
		time.Sleep(10*time.Second)
		done <- true
	}()

	for {
		select {
		case <-done:
			ticker.Stop()
			return
		case t:= <-ticker.C:
			fmt.Println("Hello !!", t)
		}
	}

	
}

func timerWatcherMain() {
	timeout:=5
	ch:=make(chan bool)

	go timer(timeout, ch)
	go watcher(timeout, ch)

	var input string

	fmt.Print("What is 725/25 ?")
	fmt.Scan(&input)

	if input == "29" {
		fmt.Println("the answer is right!")
		} else {
			fmt.Println("the answer is wrong!")
	}
}


func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\ntime out! no answer more than", timeout, "seconds")
	os.Exit(0)
}