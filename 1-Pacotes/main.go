package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo arquivo main")
	auxiliar.Escrevendo()

	erro := checkmail.ValidateFormat("ravi@gmail.com")
	fmt.Println(erro)
}