package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct {
	longradiuro string
	numero uint8
}

func main() {
	fmt.Println("Arquivo structs")
	
	var u usuario
	u.nome = "Ravi"
	u.idade = 7
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos bobos", 0}

	usuario2 := usuario{"Lorenzo", 9, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 21}
	fmt.Println(usuario3)
}