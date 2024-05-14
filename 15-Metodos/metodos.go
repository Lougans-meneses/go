package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar()  {
	// fmt.Println("Dentro do método salvar")
	fmt.Printf("Salvando os dados do Usuario %s no banco de dados \n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuário 1", 20}
	fmt.Println(usuario1)

	usuario2 := usuario{"Davi", 30}
	fmt.Println(usuario2)

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)

}