package main

import "fmt"

func somaNotas(nota, nota1 int8) int8 {
	return nota + nota1
}

func main() {

	var nota, nota1 int8

	fmt.Printf("Entre com as notas : ")
	fmt.Scanf("%d", &nota)
	fmt.Scanf("%d", &nota1)

	fmt.Println("\n\n Resultado das notas é  :", somaNotas(nota1, nota2))

}
