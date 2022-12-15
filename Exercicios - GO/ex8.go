package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func remove(texto string) {

	vogais := []string{"a", "à", "á", "ã", "â",
		"e", "è", "é", "ê", "ẽ",
		"i", "í", "ì", "ĩ", "î",
		"o", "ò", "ó", "õ", "ô",
		"u", "ù", "ú", "ũ", "û"}

	var vogaisUpper []string

	for _, e := range vogais {
		e = strings.ToUpper(e)
		vogaisUpper = append(vogaisUpper, e)
	}

	for _, e := range vogais {
		texto = strings.ReplaceAll(texto, e, "")
		for _, e := range vogaisUpper {
			texto = strings.ReplaceAll(texto, e, "")
		}
	}
	fmt.Println("Frases 0  vogais: ", texto)
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Poema: ")
	scanner.Scan()

	texto := scanner.Text()

	remove(texto)
}
