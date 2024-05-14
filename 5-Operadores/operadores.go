package main

import "fmt"

func main() {
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	// OPRECOES COM TIPOS DIFERENTES NAO REALIZA DA ERRO
	// var numero1 int16 = 10
	// var numero2 int32 = 25
	// var soma2 := numero1 + numero2
	// fmt.Println(soma2)

	
}