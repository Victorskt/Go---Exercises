package main

import "fmt"

func main() {

	fmt.Printf("Números de 1 à 100\n\n")

	var matriz [10][9]int
	a := 1

	for i := 0; i < 10; i++ {

		for j := 0; j < 9; j++ {

			matriz[i][j] = a
			a++
		}
	}

	for i := 0; i < 10; i++ {

		fmt.Println(matriz[i])
	}

	numOuro := []int{1, 1}

	for i := 0; numOuro[i] < 50; i++ {
		numOuro = append(numOuro, (numOuro[i] + numOuro[i+1]))
	}

	fmt.Printf("\n Números de 'Ouro': ")

	fmt.Println(numOuro)

}
