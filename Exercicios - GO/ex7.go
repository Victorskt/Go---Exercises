package main

import "fmt"

func varr(mediaNum, maior float64) (float64, float64) {

	var numero float64

	for i := 0; i < 10; i++ {

		fmt.Scan(&numero)
		if numero > maior {

			maior = numero
		}

		mediaNum += numero
	}

	mediaNum = mediaNum / 10

	return mediaNum, maior
}

func main() {

	var mediaNum, maior float64

	fmt.Println("Digite 10 numeros :")

	mediaNum, maior = varr(mediaNum, maior)

	fmt.Println("Média:", mediaNum)
	fmt.Println("maior numero:", maior)
}
