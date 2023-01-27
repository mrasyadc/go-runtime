package main

import (
	"fmt"
	"runtime"
	"time"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func routineMain() {
	runtime.GOMAXPROCS(2)

	startTime := time.Now()
	go print(100, "halo")
	print(100, "apa kabar")
	elapsed := time.Since(startTime)
	fmt.Printf("Binomial took %s\n", elapsed)

	var input string
	fmt.Scanln(&input)
}
