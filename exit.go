package main

import (
	"fmt"
	"os"
)

func exitMain() {
	defer fmt.Println("halo")
	os.Exit(1)
	fmt.Println("selamat datang")
}
