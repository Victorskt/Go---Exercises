package main

import "fmt"

func calcIMC(peso, altura float64) float64 {
	return peso / (altura * altura)
}

func imcTabela(imc float64) {

	fmt.Printf("Seu IMC é de %.1f: \n Classificação: ", imc)

	switch {

	case imc <= 18.5:
		fmt.Println("Abaixo do peso ideal")

	case imc >= 18.6 && imc <= 24.9:
		fmt.Println("normal")

	case imc >= 25 && imc <= 29.9:
		fmt.Println("sobrepeso")

	case imc >= 30 && imc <= 34.9:
		fmt.Println("grau I")

	case imc >= 35 && imc <= 39.9:
		fmt.Println("grau II")

	case imc >= 40:

		fmt.Println("grau III")
	}
}
func main() {

	var imc, peso, altura float64

	fmt.Printf("escreva seu Peso: ")
	fmt.Scan(&peso)

	fmt.Printf("escreva sua Altura: ")
	fmt.Scan(&altura)

	imc = calcIMC(peso, altura)

	imcTabela(imc)
}
