package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var varivel2 int = variavel1
	fmt.Println(variavel1, varivel2)

	variavel1++
	fmt.Println(variavel1)

	// PONTEIRO É REFERENCIA DE MEMORIA
	var variavel3 int
	var ponteiro *int
	
	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro)

	variavel3 = 150
	fmt.Println(variavel3, ponteiro)
}