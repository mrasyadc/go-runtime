package main

import (
	"fmt"
	"regexp"
)

func regexpMain() {
	text := "banana burger soup"
	regex, err := regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	res1 := regex.FindAllString(text, 2)
	fmt.Printf("%#v \n", res1)

	res2 := regex.FindAllString(text, -1)
	fmt.Printf("%#v \n", res2)

	regex, _ = regexp.Compile(`[a-z]+`)

	var isMatch = regex.MatchString("BHBHBBHHHASA")
	fmt.Println(isMatch)

	regex, _ = regexp.Compile(`[a-z]+`)
	text = "banana soup burger"

	str := regex.FindString(text)
	fmt.Println(str)

	text = "BANANA soup burger"
	regex, _ = regexp.Compile(`[a-z]+`)
	idx := regex.FindStringIndex(text)
	fmt.Println(idx)

	str = text[idx[0]:idx[1]]
	fmt.Println(str)

	// text1 := "banana soup burger"
	regex, _ = regexp.Compile(`[a-z]+`)

	text2 := "banana soup burger"

	stringsArr := regex.FindAllString(text2, 2)
	fmt.Println(stringsArr)


	regex, _ = regexp.Compile(`[a-z]+`)

	text2 = "banana soup burger"

	str = regex.ReplaceAllString(text2, "potato")
	fmt.Println(str)

	regex, _ = regexp.Compile(`[a-z]+`)

	text2 = "banana soup burger"

	str = regex.ReplaceAllStringFunc(text2, func (each string) string {
		if each == "burger" {
			return "potato"
		}
		return "burger"
	})
	fmt.Println(str)

	regex, _ = regexp.Compile(`[b]`)

	text2 = "banana soup burger"

	stringsArr = regex.Split(text2, -1)
	fmt.Printf("%#v \n", stringsArr)
// []string{"", "n", "n", " ", "urger soup"}
}