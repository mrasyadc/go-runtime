package main

import (
	"fmt"
	"time"
)

func timeMain() {
	time1 := time.Now()
	fmt.Printf("time1 %v\n", time1)

	time2 := time.Date(2011, 12, 24, 10, 20, 0,0, time.UTC)
	fmt.Printf("time2 %v\n", time2)
	
	
	fmt.Printf("time Now %v\n", time.Now().Year())

	year, week:= time.Now().ISOWeek()
	fmt.Printf("time Now %v%v\n", year, week)

	zone, offset:= time.Now().Zone()
	fmt.Printf("time Now %s %v\n", zone, offset)


}