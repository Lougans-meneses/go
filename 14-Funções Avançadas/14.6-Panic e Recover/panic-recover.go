package main

import "fmt"

func recuprarExecucao()  {
	if r := recover(); r != nil {
		fmt.Println("Ëxecução recuperada com  sucesso")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool  {
	media := (n1 + n2) / 12

	if media > 6 {
		return true
	} else if media < 5 {
		return false
	}

	panic("A Média é extamente 6!")
}


func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execução !")
}