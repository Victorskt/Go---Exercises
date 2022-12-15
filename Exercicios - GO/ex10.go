package main

import "fmt"

func main() {

	var resultado int
	var num int
	var i int

	fmt.Println("\n Digite um número: ")
	fmt.Scanln(&num)

	for i = 2.0; i <= num/2; i++ {

		if num%i == 0 {
			resultado++
			break
		}
	}

	if resultado == 0 {

		fmt.Print("\n", num, "é um número primo\n")
	}

	if resultado != 0 {
		fmt.Print("\n", num, " não é um número primo\n")
	}

}
