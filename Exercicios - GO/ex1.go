package main

// Arvore de Natal feita de *(Asteristicos)
import "fmt"

func main() {

	asteristico := 1

	totalDigitos := 35

	posicaoAsteristico := totalDigitos * 2

	max := (posicaoAsteristico + 1) * 3

	for i := 0; i < posicaoAsteristico; i++ {

		for count := 0; count < max; count++ {

			if count == posicaoAsteristico {

				for m := 0; m < asteristico; m++ {
					fmt.Print("*")
				}

				asteristico = asteristico + 2
				posicaoAsteristico--

			} else {

				fmt.Print(" ")
			}

		}

		fmt.Printf("\n")
	}

}
