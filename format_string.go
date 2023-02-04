package main

import "fmt"

func formatStringMain() {
	data := person{
		name: "wick",
		height: 182.5,
		age: 26,
		isGraduated: false,
		hobbies: []string{"eating", "seleeping"},
	}

	// biner %b
	fmt.Printf("%b\n", data.age)

	// karakter string unicodenya %c 
	fmt.Printf("%c\n", 1400)
	fmt.Printf("%c\n", 1325)

	// digit 10 %d
	fmt.Printf("%d\n", data.age)
	
	// scientific notation %e %E
	fmt.Printf("%e\n", data.height)
	fmt.Printf("%E\n", data.height)
	
	// decimals default 6 digit behind comma %f %F
	fmt.Printf("%f\n", data.height)
	fmt.Printf("%.9f\n", data.height)
	fmt.Printf("%.2f\n", data.height)
	fmt.Printf("%.f\n", data.height)
	
	// decimals as long as the data (can't custom digit behind comma) %g %G
	fmt.Printf("%e\n", 0.123_123_123)
	fmt.Printf("%f\n", 0.123_123_123)
	fmt.Printf("%g\n", 0.123_123_123)
	
	// octal numbers (base 8) %o
	fmt.Printf("%o\n", data.age)
	
	// pointer data %p
	fmt.Printf("%p\n", &data.age)
	
	// escape string %q
	fmt.Printf("%q\n", `"name \ height"`)
	
	// sstring %s
	fmt.Printf("%s\n", `"name \ height"`)
	
	// boolean %t
	fmt.Printf("%t\n", data.isGraduated)

	// get the variable type %T
	// string
	fmt.Printf("%T\n", data.name)

	// float64
	fmt.Printf("%T\n", data.height)

	// int32
	fmt.Printf("%T\n", data.age)

	// bool
	fmt.Printf("%T\n", data.isGraduated)

	// []string
	fmt.Printf("%T\n", data.hobbies)

	// anything anything including interface{} %v
	fmt.Printf("%v\n", data)
	
	// struct + value %+v
	fmt.Printf("%+v\n", data)
	
	// struct + value + other info %#v
	fmt.Printf("%#v\n", data)

	var otherData = struct {
    name   string
    height float64
	}{
			name:   "wick",
			height: 182.5,
	}

	// struct { name string; height float64 }{name:"wick", height:182.5}
	fmt.Printf("%#v\n", otherData)
	
	// hexadecimals (base 16) %x %X for capitals
	fmt.Printf("%x\n", data.age)

	var d = data.name

	// 7769636b
	fmt.Printf("%x%x%x%x\n", d[0], d[1], d[2], d[3])

	// 7769636b
	fmt.Printf("%x\n", d)

	// %
	fmt.Printf("%%\n")


}

type person struct {
	name string
	height float64
	age int32
	isGraduated bool
	hobbies []string
}

