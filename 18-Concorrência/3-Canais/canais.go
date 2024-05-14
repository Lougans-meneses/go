package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá mundo", canal)

	fmt.Println("Depois da função escrever começar a ser executafa")

	for {
		// <- antes da varialvel estando esperando um valor
		mensagem, aberto := <- canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}

	// for mensagem := range canal {
	// 	fmt.Println(mensagem)
	// }

	fmt.Println("Fim do programa!")
}

func escrever(texto string, canal chan string) {
	time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		// <- depos da variavel, dizendo que valor da direita esta colocando o valor na variavel a esquerda
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}