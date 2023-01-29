package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func panicRecoverMain() {
	strconvWithoutValidate()
	inputWithValidate()
	panicMain()
	recoverIife()
	recoverIifeInsideLoop()
}

func recoverIifeInsideLoop() {
	data := []string{"superman", "aquaman", "wonder woman"}

	for _, each := range data {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic occured on looping", each, "| message:", r)
				} else {
					fmt.Println("Application running perfectly")
				}
			}()
			panic("some error happen")
		}()
	}
}

func recoverIife() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error occured", r)
		} else {
			fmt.Println("Application running perfectly")
		}
	}()

	panic("something happened!")
}

func panicMain() {
	defer catch()
	var name string
	fmt.Print("Type your name:")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("halo ", name)
	} else {
		panic(err.Error())
		fmt.Println("end")
	}
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func strconvWithoutValidate() {
	var input string
	fmt.Print("Type some number:")
	fmt.Scanln(&input)

	number, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println(input, "is not a number")
		fmt.Println(err.Error())
		return
	}

	fmt.Println(number, "is number")
}

func inputWithValidate() {
	var name string
	fmt.Print("Type your name:")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("halo ", name)
	} else {
		fmt.Println(err.Error())
	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty string")
	}

	return true, nil
}
