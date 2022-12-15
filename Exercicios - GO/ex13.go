package main

import "fmt"

func main() {

	var x float64

	x = 100
	fmt.Print("\n escreva o total para a conversao : R$  ")
	fmt.Scanln(&x)

	conversor := x / 5.17
	fmt.Printf("  %0.2f", conversor)

}
