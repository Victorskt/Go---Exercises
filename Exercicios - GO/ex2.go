package main

import "fmt"

func main() {

	var ano int

	fmt.Printf(" Entre com o Ano para a conversao ! \n")
	fmt.Scanln((&ano))

	if ano <= 0 {
		fmt.Print("Ano Incorreto !  \n")

		fmt.Scanln((&ano))

	} else {

		min := ano * 525600

		seg := min * 60

		fmt.Println("\n", ano, "ano(s), equivale a : ", min, "minuto(s) e  ", seg, "segundos(s) ")
	}

}
