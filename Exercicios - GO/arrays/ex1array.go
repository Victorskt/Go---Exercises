package main

/* Escreva um algoritmo em Go
usando array que calcule a media simples de 99 valores pontos flutuantes

*/

import "fmt"

func main() {

	var arr [99]int
	soma := 0
	media := 0
	fmt.Print(len(arr))

	for i := 0; i < len(arr); i++ {
		soma += arr[i]
		media = soma / len(arr)

	}

	fmt.Println(media)
}
