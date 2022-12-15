package main

import (
	"fmt"
	"sort"
)

func main() {

	listaNomes := []string{"Joao",
		"Guilherme",
		"Diego",
		"Douglas",
		"Rodrigo",
		"Oseias",
		"Victor"}

	fmt.Println("Nomes Escolhidos:", listaNomes)

	sort.Strings(listaNomes)
	fmt.Println("Ordem Alfabética:", listaNomes)
}
