package main

import (
	"fmt"
	"runtime"
)

func channelMain() {
	runtime.GOMAXPROCS(4)

	messages := make(chan string)

	sayHelloTo := func(who string) {
		data := fmt.Sprintf("hello %s", who)
		messages <- data
	}

	// dari ketiga fungsi tersebut, goroutine yang selesai paling awal akan mengirim data lebih dulu, datanya kemudian diterima variabel message. Tanda <- jika dituliskan di sebelah kiri channel, menandakan proses penerimaan data dari channel yang di kanan, untuk disimpan ke variabel yang di kiri.

	go sayHelloTo("john wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")

	message := <-messages
	fmt.Println(message)
	message = <-messages
	fmt.Println(message)
	message = <-messages
	fmt.Println(message)

	var messages2 = make(chan string)

	for _, each := range []string{"wick", "hunt", "bourne"} {
		go func(who string) {
			data := fmt.Sprintf("hello %s", who)
			messages2 <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages2)
	}

}

func printMessage(what chan string) {
	fmt.Println(<-what)
}
