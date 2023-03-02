package main

import (
	"fmt"
	"strings"
)

func stringsMain() {
	containsString()
	hasPrefix()
	hasSuffix()
	count()
	index()
	replace()
	repeat()
	split()
	join()
	toUpper()
}

func containsString() {
	isExists := strings.Contains("john wick", "w1ck")
	fmt.Println(isExists)
}

func hasPrefix() {
	isPrefix := strings.HasPrefix("john wick", "jo")
	fmt.Println(isPrefix)
	
	isPrefix = strings.HasPrefix("john wick", "wi")
	fmt.Println(isPrefix)
}

func hasSuffix() {
	isSuffix := strings.HasSuffix("john wick", "ic")
	fmt.Println(isSuffix)
	
	isSuffix = strings.HasPrefix("john wick", "ck")
	fmt.Println(isSuffix)
}

func count() {
	howManyT := strings.Count("ethan hunt", "t")
	fmt.Println(howManyT)
}

func index() {
	index := strings.Index("ethan hunt", "n")
	fmt.Println(index)
}

func replace() {
	text := "bananana"
	find := "a"
	replaceWith := "0"

	newText := strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText)

	newText = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText)

	newText = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText)
}

func repeat() {
	str := strings.Repeat("na", 4)
	fmt.Println(str)
}

func split() {
	string1 := strings.Split("the dark knight", " ")
	fmt.Println(string1)
	string2 := strings.Split("the dark knight", "")
	fmt.Println(string2)
}

func join() {
	data := []string{"banana", "papaya", "tomato"}
	str := strings.Join(data, "-")
	fmt.Println(str)
}

func toUpper() {
	str := strings.ToUpper("eat!")
	fmt.Println(str)
}