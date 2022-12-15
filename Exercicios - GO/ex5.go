package main

import "fmt"

func medias(nota float64, nota1 float64) float64 {

	soma := nota1 / nota
	return soma
}

func main() {

	var cont int
	var notas int
	fmt.Printf("\n Quantas notas vao ser colocadas? \n")
	fmt.Scan(&cont) 
	fmt.Printf("\n")
	fmt.Printf("Digite quais notas:")
	fmt.Scan(&notas) 

	n1 := float64(cont)
	n2 := float64(notas)

	for i := 1; i < cont; i++ {
		fmt.Printf("\n______________________________________________________________________________________________________________")

		fmt.Print("\nDigite sua nota ", i, " : ")
		fmt.Println(notas)

		notas += notas
		fmt.Printf("\n______________________________________________________________________________________________________________")

	}

	fmt.Print("\n A media é: ", medias(n1, n2))

}
