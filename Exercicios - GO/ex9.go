package main

import "fmt"

func mult(n, n1, n2 int) int {

	return n * n1 * n2
}

func main() {

	var n1, n2, n3 int

	fmt.Printf(" escolha três números para multiplicar: ")

	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &n1)
	fmt.Scanf("%d", &n2)

	fmt.Println("O resultado é: ", mult(n, n1, n2))
}
