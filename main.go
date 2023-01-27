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
}
