package main

import (
	"fmt"
	"runtime"
)

func selectChannelMain() {
	runtime.GOMAXPROCS(4)
	numbers := []int{3, 4, 5, 6, 7, 5, 5, 4, 6, 7, 5, 3, 2, 5, 6, 8, 9}
	fmt.Println("numbers: ", numbers)

	ch1 := make(chan float64)
	go getAverage(numbers, ch1)

	ch2 := make(chan int)
	go getMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Avg \t: %.2f \n", avg)
		case max := <-ch2:
			fmt.Printf("Max \t: %d \n", max)
		}
	}
}

func getAverage(numbers []int, ch chan float64) {
	sum := 0
	for _, value := range numbers {
		sum += value
	}

	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	max := numbers[0]
	for _, value := range numbers {
		if value > max {
			max = value
		}
	}

	ch <- max
}
