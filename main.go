package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")

	var p1 = Person{"john wick", 22}
	fmt.Println(p1)
	p1.addAge()
	fmt.Println(p1)

	reflectMain()
	fmt.Println("===========================================")

	routineMain()
	fmt.Println("===========================================")
	channelMain()

	fmt.Println("===========================================")
	bufferedChannelMain()

	fmt.Println("===========================================")
	selectChannelMain()

	fmt.Println("===========================================")
	rangeChannelMain()

	fmt.Println("===========================================")
	// timeoutChannelMain()

	fmt.Println("===========================================")
	deferMain()

	fmt.Println("===========================================")
	pizzaDeferOrderMain()

	fmt.Println("===========================================")
	// exitMain()

	fmt.Println("===========================================")
	panicRecoverMain()

	fmt.Println("===========================================")
	formatStringMain()
	
	fmt.Println("===========================================")
	randomMain()

	fmt.Println("===========================================")
	timeMain()

	fmt.Println("===========================================")
	// timerMain()

	fmt.Println("===========================================")
	timeDurationMain()

}
