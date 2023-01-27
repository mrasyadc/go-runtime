package main

import "fmt"

type Person struct {
	name string
	age int
}

func (p *Person) addAge() {
	p.age+=1
	fmt.Println(p.age)
}