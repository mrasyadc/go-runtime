package main

import "fmt"

func deferMain() {
	defer fmt.Println("halo")
	fmt.Println("selamat datang")

	number := 3

	if number == 3 {
		fmt.Println("halo 1")
		func() {
			defer fmt.Println("halo 3")
		}()

	}

	fmt.Println("Halo 2")
}

func pizzaDeferOrderMain() {
	orderSomeFood("pizza")
	orderSomeFood("burger")
	orderSomeFood("mayones")
}

func orderSomeFood(menu string) {
	defer fmt.Println("terima kasih. silahkan tunggu")
	if menu == "pizza" {
		fmt.Print("Pilihan tepat!", " ")
		fmt.Print("Pizza di tempat kami paling enak!", "\n")
		return
	}

	fmt.Println("Pesanan anda:", menu)
}
