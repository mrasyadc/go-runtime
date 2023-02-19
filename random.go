package main

import (
	"fmt"
	"math/rand"
)

func randomMain() {
	seed:=int64(2)
	// seed:=time.Now().UTC().UnixNano()

	r:= rand.New(rand.NewSource(seed))
	
	fmt.Println("random ke-1:", r.Int())
	fmt.Println("random ke-2:", r.Int())
	fmt.Println("random ke-3:", r.Int())

	fmt.Println(r.Int())
	fmt.Println(r.Int())
	fmt.Println(r.Int())
	
	fmt.Println(r.Intn(100))
	fmt.Println(r.Intn(100))
	fmt.Println(r.Intn(100))

	fmt.Println(randomString(1000, 2))

}



func randomString(length int, seed int64) string {
	r:= rand.New(rand.NewSource(seed))
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b:=make([]rune, length)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}

	return string(b)
}
