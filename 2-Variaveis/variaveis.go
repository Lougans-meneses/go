package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1"
	// atribuicao ppor inferencia 
	variavel2 := "Variavel 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 = "Variavel 3"
		variavel4 = "Variavel 4"
	)

	fmt.Println(variavel3, variavel4)

	// multiplas atribuicoes por inferencia de tipo
	variavel5, variavel6 := "variavel 5", "variavel 6"
	fmt.Println(variavel5, variavel6)

	// constantes imutaveis que nao podem ser reatribuidas
	const constante1 string = "constante 1"
	fmt.Println(constante1)

	// Invertendo os valores 
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}