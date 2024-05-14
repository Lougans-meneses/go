package main

import (
	"errors"
	"fmt"
)

func main() {
	var num int64 = -1000000
	fmt.Println(num)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// alias 
	// int32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	// byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.26
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12300000000.23
	fmt.Println(numeroReal2)

	numeroReal3 := 123456.67
	fmt.Println(numeroReal3)

	// Strings

	var str string = "texto"
	fmt.Println(str)

	str2 := "texto2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	// nao declarada assume 0
	var booleno1 bool 
	fmt.Println(booleno1)

	// nill para vareavel error sem declaracao
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}