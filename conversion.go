package main

import (
	"fmt"
	"strconv"
)

func conversionMain() {
	stringToInt()
	intToString()
	parseInt()
	parseIntBinary()
	formatInt()
	parseFloat()
	formatFloat()
	parseBool()
	formatBool()
	casting()
	stringToByte()
	byteToString()
	assertionInterface()
}

func stringToInt() {
	str := "123"
	num, err := strconv.Atoi(str)

	if err == nil {
		fmt.Println(num)
	}
}

func intToString() {
	num := 123
	str := strconv.Itoa(num)
	fmt.Println(str)
}

func parseInt() {
	str := "123"
	num, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		fmt.Println(num)
	}
}

func parseIntBinary() {
	str:= "1010"
	num, err := strconv.ParseInt(str, 2, 8)

	if err == nil {
		fmt.Println(num)
	}
}

func formatInt() {
	num := int64(24)
	str := strconv.FormatInt(num, 2)

	fmt.Println(str)
}

func parseFloat() {
	str:= "12.12"
	num, err := strconv.ParseFloat(str, 64)

	if err == nil {
		fmt.Println(num)
	}
}

func formatFloat() {
	num := float64(12.12)
	str := strconv.FormatFloat(num, 'f', 6, 64)

	fmt.Println(str)
}

func parseBool() {
	str:= "true"
	bul, err := strconv.ParseBool(str)

	if err == nil {
		fmt.Println(bul)
	}
}

func formatBool() {
	bul := true
	str := strconv.FormatBool(bul)

	fmt.Println(str)
}

func casting() {
	a := float64(24)
	b := int32(24.00)

	c := int32(a)
	d := float64(b)

	fmt.Println(a, b, c, d)
}

func stringToByte() {
	text := "Halo"
	b := []byte(text)

	fmt.Printf("%d %d %d %d", b[0], b[1],b[2],b[3],)
}

func byteToString() {
	textByte := []byte{72, 97,108, 111}

	text := string(textByte)

	fmt.Printf("\n%s\n", text)

	c := int32('h')
	fmt.Println(c) // 104

	d := string(c)

	fmt.Println(d)
}

func assertionInterface() {
	data := map[string]interface{} {
		"nama": "john wick",
		"grade": 2,
		"height": 156.5,
		"isMale": true,
		"hobbies": []string{"eat", "sleep", "code", "repeat"},
	}

	fmt.Println(data["nama"].(string))
	fmt.Println(data["grade"].(int))
	fmt.Println(data["height"].(float64))
	fmt.Println(data["isMale"].(bool))
	fmt.Println(data["hobbies"].([]string))

	for _, val := range data {
		switch val.(type) {
		case string:
			fmt.Println(val.(string))
		
		case int:
			fmt.Println(val.(int))

		case float64:
			fmt.Println(val.(float64))

		case bool:
			fmt.Println(val.(bool))

		case []string:
			fmt.Println(val.([]string))

		default:
			fmt.Println(val.(int))
		}
	}
}
