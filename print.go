package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func routineMain() {
	// runtime.GOMAXPROCS(4)
	fmt.Println("runtime cpu: ", runtime.NumCPU())

	go print(100, "halo")
	print(100, "apa kabar")

	// Karena pengiriman dan penerimaan data lewat channel bersifat blocking, tidak perlu memanfaatkan sifat blocking dari fungsi sejenis fmt.Scanln() atau lainnya, untuk mengantisipasi goroutine utama main selesai sebelum ketiga goroutine di atas selesai.

	// oleh karena itu digunakan fmt.Scanln() yang blocking untuk mengantisipasi goroutine utama main selesai sebelum ini selesai

	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
