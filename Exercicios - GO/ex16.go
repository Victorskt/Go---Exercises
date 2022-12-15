package main

import (
	"fmt"
	"time"
)

func main() {

	var dia int
	var mes int
	var ano int

	fmt.Println("Digite a Data  dia/mês/ano,  no Padrão dd/MM/yyyy")
	fmt.Scan(&dia)
	fmt.Scan(&mes)
	fmt.Scan(&ano)

	resultado := time.Date(ano, time.Month(mes), dia, 00, 00, 00, 00, time.Local)

	fmt.Println("Você Nasceu em uma:")

	if resultado.Weekday().String() == "Sunday" {

		fmt.Printf("Domingo")

	}

	if resultado.Weekday().String() == "Monday" {

		fmt.Println("Segunda-feira")

	}

	if resultado.Weekday().String() == "Tuesday" {

		fmt.Println("Terça-feira")

	}

	if resultado.Weekday().String() == "Wednesday" {

		fmt.Println("Quarta-feira")

	}

	if resultado.Weekday().String() == "Thursday" {

		fmt.Println("Quinta-feira")

	}

	if resultado.Weekday().String() == "Friday" {

		fmt.Println("Sexta-feira")

	}

	if resultado.Weekday().String() == "Saturday" {

		fmt.Println("Sábado")

	}
}
