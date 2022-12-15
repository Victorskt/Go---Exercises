package main

import "fmt"

func main() {

	fmt.Printf("\nDigite a nota: ")
	var nota int

	fmt.Scanln(&nota)

	if nota >= 0 && nota < 7 {

		fmt.Printf("\nA avaliação desse cliente é: detrator")

	}

	if nota > 6 && nota < 9 {

		fmt.Printf("\nA avaliação desse cliente é: neutro")

	}

	if nota > 8 && nota <= 10 {

		fmt.Printf("\nA avaliação desse cliente é: promotor ")

	}

	if nota < 0 || nota > 10 {

		fmt.Printf("\nAvaliação Fora da Métrica\n")

	}
}
