package main

import "fmt"

func main() {
	fmt.Println("hello world")
	printFromPrint()
	printFromPrint()

	var p1 = Person{"john eick", 22}
	fmt.Println(p1)
	p1.addAge()
	fmt.Println(p1)
}