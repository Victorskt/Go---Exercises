package main

import "fmt"

func main() {

	var x float32

	fmt.Printf(" \ndigite a temperatura em Celcius:  ")
	fmt.Scanln(&x)

	fahr := (x * 1.8) + 32

	fmt.Print("\n O resultado é:  ", fahr)

}
