package main

import "fmt"

func main() {

	var num int

	fmt.Printf("Escreva um numero (segundos) que será convertido para (minuntos) :\n")
	fmt.Scanln(&num)

	if num < 60 {

		fmt.Println("vale a :", num, "segundos")

	} else {

		min := num / 60
		seg := num % 60
		fmt.Println("\n", num, "segundos, vale a :", min, "minuto(s) e ", seg, "segundo(s)")

	}

}
