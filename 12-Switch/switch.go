package main

import "fmt"

func diaDaSemana(numero int) string {
	var diaDaSemana string
	switch {
	case numero == 1:
		diaDaSemana =  "Domingo"
		fallthrough // ESSE CARA FAZ A CONDICAO CAIR NA PROXIMA FASE
	case numero ==2:
		diaDaSemana = "Segunda-feira"
	case numero == 3:
		diaDaSemana = "Terça-feira"
	case numero == 4:
		diaDaSemana = "Quarta-feira"
	case numero == 5:
		diaDaSemana = "Quinta-feira"
	case numero == 6:
		diaDaSemana = "Sexta-feira"
	case numero == 7:
		diaDaSemana = "Sabado"
	case numero == 8:
		diaDaSemana = "Domingo"
	default:
		diaDaSemana = "Número inválido"
	}
	return diaDaSemana
}

func diaDaSemana2(numero int) string {
	switch  {
	case numero == 1:
		return "Domingo"
	}
	return ""
}

func main() {
	dia := diaDaSemana(2)
	fmt.Println(dia)


	fmt.Println("--------")
	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)

}