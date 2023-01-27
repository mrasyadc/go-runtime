package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
	Age   int
}

func (s *student) SetNameAge(name string, age int) {
	s.Name = name
	s.Age += age
}

func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama      :", reflectType.Field(i).Name)
		fmt.Println("tipe data :", reflectType.Field(i).Type)
		fmt.Println("nilai     :", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func reflectMain() {
	fmt.Println("==================================================")
	var number = 23
	var reflectValue = reflect.ValueOf(number)
	fmt.Println("tipe  variabel :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel :", reflectValue.Int())
	}

	// using interface to print
	fmt.Println("using interface to print")
	fmt.Println("tipe  variabel :", reflectValue.Type())
	fmt.Println("nilai variabel :", reflectValue.Interface())

	// casting interface
	var nilai = reflectValue.Interface().(int)
	fmt.Printf("type of nilai is %T\n", nilai)

	fmt.Println("==================================================")
	// akses informasi property variable objek
	var s1 = &student{Name: "wick", Grade: 2, Age: 3}
	s1.getPropertyInfo()

	// akses informasi method variable objek
	reflectValue = reflect.ValueOf(s1)
	var method = reflectValue.MethodByName("SetNameAge")

	method.Call([]reflect.Value{
		reflect.ValueOf("john"),
		reflect.ValueOf(7),
	})

	fmt.Println("nama :", s1.Name)
	fmt.Println("nama :", s1.Age)
}
