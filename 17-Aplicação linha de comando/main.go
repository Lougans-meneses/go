package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de começar")

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}