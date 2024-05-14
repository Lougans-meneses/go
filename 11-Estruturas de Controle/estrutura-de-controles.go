package main

import "fmt"

func main() {
	fmt.Println("Estrutura de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// INFINITY INICIANDO A VARIAVEL JA NO IF
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else if numero < -10{
		fmt.Println("Número é menor que -10")
	} else {
		fmt.Println("Número é menor que zero")
	}

	fmt.Println()
}