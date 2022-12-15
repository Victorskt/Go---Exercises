package main

import "fmt"

func somar(nota int, nota1 int) int {

	soma := nota + nota1
	return soma
}

func main() {

	var num1 int
	var num2 int

	fmt.Printf("\ndigite o valor da primeira nota \n")
	fmt.Scanln(&num1)

	fmt.Printf("\n\n")
	fmt.Printf("\ndigite o valor da segunda nota \n")
	fmt.Scanln(&num2)

	fmt.Printf("\n\n")
	fmt.Println("A soma das notas é : ", somar(num1, num2))

}
