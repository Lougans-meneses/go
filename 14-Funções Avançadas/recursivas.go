package main

import "fmt"

func fibonacci(posicao uint) uint  {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções recursivas")

	posicao := uint(12)

	for i := uint(1); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}

	fmt.Println(fibonacci(posicao))
}